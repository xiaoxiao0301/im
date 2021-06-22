package controller

import (
	"hello/services"
	"hello/util"
	"net/http"
	"strconv"
)

var groupServices services.GroupService
var communityServices services.CommunityService

// Joincommunity 加入群聊
func Joincommunity(response http.ResponseWriter, request *http.Request) {
	//	dstid:dstobj,userid:userId()
	dstIdStr := request.FormValue("dstid")
	userIdStr := request.FormValue("userid")
	dstId, _ := strconv.Atoi(dstIdStr)
	userId, _ := strconv.Atoi(userIdStr)
	memo := "申请加入群聊"
	isExists := communityServices.IsExistsCommunity(int64(dstId))
	if !isExists {
		util.RespFail(response, "您要加入的群聊不存在")
		return
	}

	group, err := groupServices.UserAddGroup(int64(userId), int64(dstId), memo)
	if err != nil {
		util.RespFail(response, err.Error())
	} else {
		util.RespSuccess(response, group, "加入成功")
	}
}
