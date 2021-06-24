package services

import (
	"hello/config"
	"hello/model"
	"log"
	"math"
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

// GetUserMessageList 查找登录用户与当前用户的聊天信息 cmd 10 私聊 11 群聊
func (this *MessageService) GetUserMessageList(cmd int, page int, userId, distId int64) (messages []model.Message, totalPage int) {

	messages = make([]model.Message, 0)
	// 获取总记录条数
	countRows, err := config.GetDbEngine().Where("user_id = ? and dst_id = ? and cmd = ?", userId, distId, cmd).Count(&model.Message{})
	if err != nil {
		log.Println(err.Error())
	}
	if countRows <= 0 {
		return messages, 0
	}
	// 总页数
	totalPage = int(math.Ceil(float64(countRows) / float64(config.PAGE_SIZE)))
	if page <= 0 {
		page = 1
	}
	if page >= totalPage {
		page = totalPage
	}

	config.GetDbEngine().Where("user_id = ? and dst_id = ? and cmd = ?", userId, distId, cmd).
		Limit(config.PAGE_SIZE, config.PAGE_SIZE*(page-1)).Find(&messages)
	return messages, totalPage
}
