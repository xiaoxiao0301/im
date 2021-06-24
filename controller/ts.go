package controller

import (
	"hello/config"
	"hello/model"
	"hello/util"
	"net/http"
)

func T(response http.ResponseWriter, request *http.Request) {
	config.GetDbEngine().InsertOne(&model.TestTime{
		Name: "第一条哈",
	})
	util.RespSuccess(response, nil, "ok")
}

func L(response http.ResponseWriter, request *http.Request) {
	t := model.TestTime{}
	config.GetDbEngine().ID(1).Get(&t)
	util.RespSuccess(response, t, "ok")
}
