package services

import (
	"errors"
	"hello/config"
	"hello/model"
	"log"
	"math"
)

type ContactService struct {
}

// AddFriend 添加好友
func (this *ContactService) AddFriend(userId int, dstId int, memo string) error {
	if userId == dstId {
		return errors.New("不能添加自己为好友")
	}
	tmpContast := model.Contact{}
	// 判断当前用户是否已经添加过了
	oks, err := config.GetDbEngine().Where("user_id = ?", userId).And("dst_id = ?", dstId).Get(&tmpContast)
	if err != nil {
		return err
	}
	if oks {
		return errors.New("该用户已经添加过啦")
	}
	// 开启事务添加好友
	session := config.GetDbEngine().NewSession()
	defer session.Close()
	session.Begin()

	_, err1 := config.GetDbEngine().InsertOne(model.Contact{
		UserId: int64(userId),
		DstId:  int64(dstId),
		Type:   model.CONTACT_TYPE_USER,
		Memo:   memo,
	})

	_, err2 := config.GetDbEngine().InsertOne(model.Contact{
		UserId: int64(dstId),
		DstId:  int64(userId),
		Type:   model.CONTACT_TYPE_USER,
		Memo:   memo,
	})

	if err1 == nil && err2 == nil {
		// 提交事务
		session.Commit()
		return nil
	} else {
		// 任何一个出错，回滚事务
		session.Rollback()
		if err1 != nil {
			return err1
		} else {
			return err2
		}
	}
}

// SearchFriend 获取好友列表
func (this *ContactService) SearchFriend(userId int64, page int) (comms []model.User, totalPages int) {
	// 最后的返回结果
	comms = make([]model.User, 0)

	// 获取当前用户的总数
	contact := model.Contact{}
	countRows, errs := config.GetDbEngine().Where("user_id = ? and type = ?", userId, model.CONTACT_TYPE_USER).Count(&contact)
	if errs != nil {
		log.Fatalln(errs.Error())
	}
	// 没有好友情况
	if countRows <= 0 {
		return comms, 0
	}
	// 总页数
	totalPages = int(math.Ceil(float64(countRows) / float64(config.PAGE_SIZE)))
	if page >= totalPages {
		page = totalPages
	}
	if page <= 0 {
		page = 1
	}

	// mysql-limit分页 0, 10 10, 10
	userContacts := make([]model.Contact, 0)
	config.GetDbEngine().Where("user_id = ? and type = ?", userId, model.CONTACT_TYPE_USER).Limit(config.PAGE_SIZE, (page-1)*config.PAGE_SIZE).Find(&userContacts)

	// 获取好友的所有id
	dstIds := make([]int64, 0)
	for _, v := range userContacts {
		dstIds = append(dstIds, v.DstId)
	}

	config.GetDbEngine().In("id", dstIds).Find(&comms)

	return comms, totalPages
}
