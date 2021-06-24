package main

import (
	"fmt"
	"hello/controller"
	"log"
	"net/http"
	"text/template"
)

// RegisterView 解析模板，自动导入
func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		// 打印并直接退出
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Println(tplname)
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {

	// 用户登录
	http.HandleFunc("/user/login", controller.UserLogin)
	// 用户注册
	http.HandleFunc("/user/register", controller.UserRegister)
	// 添加好友
	http.HandleFunc("/contact/addfriend", controller.AddFriend)
	// 获取好友列表
	http.HandleFunc("/contact/friend-list", controller.FriendList)
	// 查找好友
	http.HandleFunc("/user/find", controller.UserFind)
	// 创建群
	http.HandleFunc("/contact/createcommunity", controller.CreateCommunity)
	// 加入群聊
	http.HandleFunc("/user/joincommunity", controller.Joincommunity)
	// 获取群聊列表
	http.HandleFunc("/contact/community", controller.CommunityList)
	// 文件上传
	http.HandleFunc("/attach/upload", controller.UploadFileHelper)
	// 聊天
	http.HandleFunc("/chat", controller.Chat)
	// 获取聊天记录列表
	http.HandleFunc("/message/list", controller.MessageList)
	// t-测试自定义的添加时间
	http.HandleFunc("/test", controller.T)
	// l-列表输出格式时间
	http.HandleFunc("/test-list", controller.L)

	// 静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))
	// 提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// 文件上传解析
	http.Handle("/upload/", http.FileServer(http.Dir(".")))

	/*http.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
		tpl, err := template.ParseFiles("view/user/login.html")
		if err != nil {
			// 打印并直接退出
			log.Fatal(err.Error())
		}
		tpl.ExecuteTemplate(writer, "/user/login.shtml", nil)
	})*/

	RegisterView()

	http.ListenAndServe(":8080", nil)
}
