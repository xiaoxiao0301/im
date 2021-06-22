package services

import (
	"errors"
	"hello/config"
	"hello/model"
	"log"
	"math"
)

type CommunityService struct {
}

// CreateCommunity 创建群聊
func (this *CommunityService) CreateCommunity(userId int64, cate int, name, icon, memo string) (comm model.Community, err error) {
	comm = model.Community{}
	// ok 数据库有记录返回 true 无记录返回 false
	ok, err := config.GetDbEngine().Where("name = ? and user_id = ? and cate = ?", name, userId, cate).Get(&comm)
	if err != nil {
		return comm, err
	}
	if ok {
		return comm, errors.New("群聊已经存在了")
	}

	comm.UserId = userId
	comm.Name = name
	comm.Icon = icon
	comm.Cate = this.getCommunityType(cate)
	comm.Memo = memo
	comm.Number = 1

	_, err = config.GetDbEngine().InsertOne(&comm)
	return comm, err
}

// CommunityList 群聊列表
func (this *CommunityService) CommunityList(userId int64, page int) (comm []model.Community, totalPage int) {
	comm = make([]model.Community, 0)
	community := model.Community{}
	rowCounts, err := config.GetDbEngine().Where("user_id = ?", userId).Count(&community)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// 没有创建群聊或加入群聊
	if rowCounts <= 0 {
		return comm, 0
	}

	totalPage = int(math.Ceil(float64(rowCounts) / float64(config.PAGE_SIZE)))
	if page >= totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	config.GetDbEngine().Where("user_id = ?", userId).Limit(config.PAGE_SIZE, (page-1)*config.PAGE_SIZE).Find(&comm)
	return comm, totalPage
}

// 获取群类型
func (this *CommunityService) getCommunityType(types int) int {
	var cate int
	switch types {
	case 1:
		cate = model.COMMUNITY_CATE_COM
	case 2:
		cate = model.COMMUNITY_CATE_HOBBY
	case 3:
		cate = model.COMMUNITY_CATE_INDUSTRY
	case 4:
		cate = model.COMMUNITY_CATE_LIFE
	case 5:
		cate = model.COMMUNITY_CATE_STUDY
	}
	return cate
}
