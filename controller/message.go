package controller

import (
	"hello/util"
	"net/http"
	"strconv"
)

// MessageList 获取聊天信息
func MessageList(response http.ResponseWriter, request *http.Request) {
	userid := request.FormValue("userid")
	dstid := request.FormValue("dstid")
	pageStr := request.FormValue("page")
	cmdStr := request.FormValue("cmd") // 聊天类型 10 私聊 11 群聊
	userId, _ := strconv.ParseInt(userid, 10, 64)
	dstId, _ := strconv.ParseInt(dstid, 10, 64)
	page, _ := strconv.Atoi(pageStr)
	cmd, _ := strconv.Atoi(cmdStr)
	message, totalPage := messageServices.GetUserMessageList(cmd, page, userId, dstId)
	util.RespSuccessList(response, message, totalPage, page, "获取成功")
}
