package controller

import (
	"fmt"
	"hello/model"
	"hello/services"
	"hello/util"
	"math/rand"
	"net/http"
)

var userService services.UserService

// UserRegister 用户注册
func UserRegister(response http.ResponseWriter, request *http.Request) {
	mobile := request.FormValue("mobile")
	password := request.FormValue("password")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW
	user, err := userService.RegisterUser(mobile, password, nickname, avatar, sex)
	if err != nil {
		util.RespFail(response, err.Error())
	} else {
		util.RespSuccess(response, user, "注册成功!")
	}
}

// UserLogin 用户登录
func UserLogin(response http.ResponseWriter, request *http.Request) {
	mobile := request.FormValue("mobile")
	password := request.FormValue("password")

	user, err := userService.LoginUser(mobile, password)
	if err != nil {
		util.RespFail(response, err.Error())
	} else {
		util.RespSuccess(response, user, "登陆成功!")
	}
}
