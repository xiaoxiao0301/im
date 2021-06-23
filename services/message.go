package services

import (
	"hello/config"
	"hello/model"
)

type MessageService struct {
}

// SaveChatMessage 存储消息
func (this *MessageService) SaveChatMessage(userid int64, dstid int64, cmd int, media int,
	content string, pic string, url string, memo string, amount int) {
	//{"dstid":1,"cmd":10,"userid":2,"media":1,"content":"你好，1"}
	message := model.Message{
		UserId:  userid,
		DstId:   dstid,
		Cmd:     cmd,
		Media:   media,
		Content: content,
		Pic:     pic,
		Url:     url,
		Memo:    memo,
		Amount:  amount,
	}

	config.GetDbEngine().InsertOne(&message)

}
