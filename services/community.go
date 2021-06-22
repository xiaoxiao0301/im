package services

import (
	"errors"
	"hello/config"
	"hello/model"
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

func (this *CommunityService) GetCommunityInfo(communityId int64) (comm model.Community) {
	config.GetDbEngine().Where("id = ?", communityId).Get(&comm)
	return comm
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
