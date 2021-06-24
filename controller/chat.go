package controller

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"hello/model"
	"hello/services"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// Node
type Node struct {
	Conn      *websocket.Conn //websocket连接
	DataQueue chan []byte     // 并行转串行
	GroupSets set.Interface   // set包，用来计算 交集，并集，差集  go get -u gopkg.in/fatih/set.v0
}

// 用户id和node绑定
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwlocker sync.RWMutex

// 存储局域网广播消息
var udpSendScan chan []byte = make(chan []byte, 1024)

var userServices services.UserService
var messageServices services.MessageService

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// Chat 聊天
func Chat(response http.ResponseWriter, request *http.Request) {
	// 身份校验 http://localhost:8080/chat?id=1&token=23B9D063E65095EC47D634198BF902BD
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userid, _ := strconv.ParseInt(id, 10, 64)
	// websocket 使用 https://github.com/gorilla/websocket
	// websocket 是在http的基础上升级，会附带upgrade请求
	check := checkToken(userid, token)
	// conn是连接
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 用于判断对当前的请求时拦截还是放行，这里当用户的token错误的时候拦截请求
		CheckOrigin: func(request *http.Request) bool {
			return check
		},
	}).Upgrade(response, request, response.Header())
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),   // 存储的消息体
		GroupSets: set.New(set.ThreadSafe), //线程安全, 存储所有用的群聊id，已经加入的，会自动去重，消息群聊使用
	}
	// 获取用户全部群ID ， 目的是通过判断集合中有消息的群聊id后发送
	groupIds := groupServices.SearchCommunityIds(userid)
	for _, v := range groupIds {
		// set包使用，添加元素
		node.GroupSets.Add(v)
	}
	// 加锁，将用户和node绑定到一起
	rwlocker.Lock()
	clientMap[userid] = node
	rwlocker.Unlock()
	go sendProc(node)
	go recvProc(node)
	log.Printf("%d is connected\n", userid)
}

// ws 发送
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data) // 第一个参数是消息类型，text类型 ，第二个参数是消息内容
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// ws 接收
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("recv<=%s\n", string(data))
		// 把消息广播到局域网
		broadMsg(data)
		//dispathMessage(data)
	}
}

// udp 发送
func udpSendProc() {
	log.Println("into start udp send proc")
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	if err != nil {
		log.Println("conn upd server err:", err.Error())
		return
	}
	defer conn.Close()
	// conn发送数据
	for true {
		select {
		case data := <-udpSendScan:
			_, err = conn.Write(data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// udp 接收
func udpRecvProc() {
	log.Println("into start upd recv proc")
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		log.Println("listen upd fail err ", err.Error())
		return
	}
	for true {
		buff := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Println(err.Error())
			return
		}
		// 数据处理
		dispathMessage(buff[:n])
	}
	log.Println("stop udp recv proc")
}

// 根据消息类型处理
func dispathMessage(data []byte) {
	message := model.Message{}
	// 解析消息
	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 将用户的聊天信息存储起来
	messageServices.SaveChatMessage(message.UserId, message.DstId, message.Cmd, message.Media, message.Content,
		message.Pic, message.Url, message.Memo, message.Amount)
	switch message.Cmd {
	case model.MESSAGE_CMD_SINGLE:
		// 单聊消息
		log.Printf("c2cmsg %d=>%d\n%s\n", message.UserId, message.DstId, string(data))
		sendMsg(message.DstId, data)
	case model.MESSAGE_CMD_GROUP:
		// 群聊消息
		for _, v := range clientMap {
			if v.GroupSets.Has(message.DstId) {
				v.DataQueue <- data
			}
		}
		log.Printf("c2gmsg %d=>%d\n%s\n", message.UserId, message.DstId, string(data))
	case model.MESSAGE_CMD_HEART:
		// 心跳信息，啥都不做
	default:
		//
	}
}

func sendMsg(userId int64, data []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- data
	}
}

// 将消息发送到局域网中
func broadMsg(data []byte) {
	udpSendScan <- data
}

// 用户身份令牌校验
func checkToken(userid int64, token string) bool {
	return userServices.FindUser(userid).Token == token
}

// AddGroupIdToUserGroupSet 用户新创建的群id加入到用户的群组set中
func AddGroupIdToUserGroupSet(userId, groupId int64) {
	rwlocker.Lock()
	node, ok := clientMap[userId] // 获取node
	if ok {
		node.GroupSets.Add(groupId)
	}
	rwlocker.Unlock()
}
