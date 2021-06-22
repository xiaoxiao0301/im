package controller

import (
	"hello/services"
	"hello/util"
	"net/http"
	"strconv"
)

var contactService services.ContactService

// AddFriend 添加好友
func AddFriend(response http.ResponseWriter, request *http.Request) {
	// dstid:dstobj,userid: user.id,pic:user.avatar,content:user.nickname,memo: "请求加你为好友"
	userIdStr := request.FormValue("userid")
	dstIdStr := request.FormValue("dstid")
	memo := request.FormValue("memo")
	userId, _ := strconv.Atoi(userIdStr)
	dstId, _ := strconv.Atoi(dstIdStr)
	err := contactService.AddFriend(userId, dstId, memo)
	if err != nil {
		util.RespFail(response, err.Error())
	} else {
		util.RespSuccess(response, nil, "好友添加成功")
	}
}

// FriendList 好友列表
func FriendList(response http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userid")
	userId, _ := strconv.Atoi(userIdStr)
	pageStr := request.FormValue("page")
	page, _ := strconv.Atoi(pageStr)
	users, total := contactService.SearchFriend(int64(userId), page)
	util.RespSuccessList(response, users, total, page, "获取成功")
}
