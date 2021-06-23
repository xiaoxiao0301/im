package model

import "time"

type Message struct {
	Id        int64     `xorm:"bigint(20) pk autoincr" json:"id"`
	UserId    int64     `xorm:"bigint(20) comment('发送用户id')" json:"userid,omitempty"`
	DstId     int64     `xorm:"bigint(20) comment('接收id，用户id或者群聊id')" json:"dstid,omitempty"`
	Cmd       int       `xorm:"int(11) comment('消息类别，单聊还是群聊')" json:"cmd,omitempty"`
	Media     int       `xorm:"int(11) comment('消息样式，文本，图片，语音，视频')" json:"media,omitempty"`
	Content   string    `xorm:"varchar(255) comment('消息内容')" json:"content,omitempty"`
	Pic       string    `xorm:"varchar(255) comment('预览图片')" json:"pic,omitempty"`
	Url       string    `xorm:"varchar(255) comment('服务的URL')" json:"url,omitempty"`
	Memo      string    `xorm:"varchar(255) comment('简单描述')" json:"memo,omitempty"`
	Amount    int       `xorm:"int(11) comment('金额，扩展字段')" json:"amount,omitempty"`
	CreatedAt time.Time `xorm:"datetime created" json:"create_at"`
	UpdatedAt time.Time `xorm:"datetime updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"datetime deleted" json:"deleted_at"`
}

// 消息类别
const (
	MESSAGE_CMD_SINGLE = 10 // 单聊
	MESSAGE_CMD_GROUP  = 11 // 群聊
	MESSAGE_CMD_HEART  = 0  // 心跳监测信息
)

// 消息样式
const (
	MESSAGE_MEDIA_TEXT       = 1 // 文本消息
	MESSAGE_MEDIA_NEWS       = 2 // 新闻样式，类比图文消息
	MESSAGE_MEDIA_VOICE      = 3 // 语音样式
	MESSAGE_MEDIA_IMAGE      = 4 // 图片样式包含emoj
	MESSAGE_MEDIA_REDPACKAGR = 5 // 红包样式
	//MESSAGE_MEDIA_EMOJ   = 6 // emoj表情样式
	MESSAGE_MEDIA_LINK    = 7   // 超链接样式
	MESSAGE_MEDIA_VIDEO   = 8   // 视频样式
	MESSAGE_MEDIA_CONCAT  = 9   // 名片样式
	MESSAGE_MEDIA_DEFAULT = 100 // 自定义
)

// 文本消息体 {id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
// 图片，音频，视频消息体 {id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}

/**
消息发送结构体
1、MEDIA_TYPE_TEXT
{id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}

2、MEDIA_TYPE_News
{id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}

3、MEDIA_TYPE_VOICE，amount单位秒
{id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}

4、MEDIA_TYPE_IMG
{id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}

5、MEDIA_TYPE_REDPACKAGR //红包amount 单位分
{id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}

6、MEDIA_TYPE_EMOJ 6
{id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}

7、MEDIA_TYPE_Link 6
{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}

8、MEDIA_TYPE_VIDEO 8
{id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}

9、MEDIA_TYPE_CONTACT 9
{id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}

*/
