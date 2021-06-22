package services

import (
	"hello/config"
	"hello/model"
	"log"
	"math"
)

type GroupService struct {
}

type Result struct {
	UserId      int64  `json:"user_id"`
	CommunityId int64  `json:"community_id"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Memo        string `json:"memo"`
}

// CommunityList 群聊列表
func (this *GroupService) CommunityList(userId int64, page int) (result []Result, totalPage int) {
	result = make([]Result, 0)
	rowCounts, err := config.GetDbEngine().Where("user_id = ? and type = ?", userId, model.CONTACT_TYPE_COMMUNITY).Count(&model.Group{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	// 没有创建群聊或加入群聊
	if rowCounts <= 0 {
		return result, 0
	}

	totalPage = int(math.Ceil(float64(rowCounts) / float64(config.PAGE_SIZE)))
	if page >= totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	group := make([]model.Group, 0)
	config.GetDbEngine().Where("user_id = ? and type = ?", userId, model.CONTACT_TYPE_COMMUNITY).Limit(config.PAGE_SIZE, (page-1)*config.PAGE_SIZE).Find(&group)
	// 获取对应的群聊详细信息
	for _, v := range group {
		var communityService CommunityService
		commu := communityService.GetCommunityInfo(v.DstId)
		result = append(result, Result{
			UserId:      v.Id,
			CommunityId: v.DstId,
			Icon:        commu.Icon,
			Name:        commu.Name,
			Memo:        commu.Memo,
		})
	}

	return result, totalPage
}

// AddGroup 创建群聊成功后需要添加一条记录到群列表中
func (this *GroupService) AddGroup(userid, dstid int64) {
	group := model.Group{
		UserId: userid,
		DstId:  dstid,
		Type:   model.CONTACT_TYPE_COMMUNITY,
		Memo:   "",
	}
	config.GetDbEngine().InsertOne(&group)

}

func (this *GroupService) AddCommunity(userId int64, communityId int64, memo string) (group model.Group, err error) {
	// 不能重复加入

	// 要添加的群是否存在

	return group, nil
}
