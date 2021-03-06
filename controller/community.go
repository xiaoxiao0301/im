package controller

import (
	"hello/services"
	"hello/util"
	"net/http"
	"strconv"
)

var communityService services.CommunityService
var groupService services.GroupService

// CreateCommunity 创建群聊
func CreateCommunity(response http.ResponseWriter, request *http.Request) {
	// name: 测试  cate: 1 memo: 测试  icon: /asset/images/community.png  userid: 1
	name := request.FormValue("name")
	icon := request.FormValue("icon")
	memo := request.FormValue("memo")
	cate, _ := strconv.Atoi(request.FormValue("cate"))
	userId, _ := strconv.Atoi(request.FormValue("userid"))

	comm, err := communityService.CreateCommunity(int64(userId), cate, name, icon, memo)
	if err != nil {
		util.RespFail(response, err.Error())
	} else {
		groupService.AddGroup(int64(userId), int64(comm.Id))
		AddGroupIdToUserGroupSet(int64(userId), comm.Id) // 创建群聊后将群聊id加入用户群组的set中
		util.RespSuccess(response, comm, "创建群聊成功")
	}
}

// CommunityList 群聊列表
func CommunityList(response http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userid")
	pageStr := request.FormValue("page")
	userid, _ := strconv.Atoi(userIdStr)
	page, _ := strconv.Atoi(pageStr)
	list, total := groupService.CommunityList(int64(userid), page)
	util.RespSuccessList(response, list, total, page, "获取成功")
}
