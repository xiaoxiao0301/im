package controller

import (
	"hello/services"
	"hello/util"
	"net/http"
	"strconv"
)

var userServices services.UserService

// Chat 聊天
func Chat(response http.ResponseWriter, request *http.Request) {
	// 身份校验 http://localhost:8080/chat?id=1&token=23B9D063E65095EC47D634198BF902BD
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userid, _ := strconv.Atoi(id)
	check := checkToken(int64(userid), token)
	if !check {
		util.RespFail(response, "令牌错误，请重新登录")
		return
	}

	// websocket 使用 https://github.com/gorilla/websocket

}

// 用户身份令牌校验
func checkToken(userid int64, token string) bool {
	return userServices.FindUser(userid).Token == token
}
