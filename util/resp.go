package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` //第二个是为空时不显示
}

func RespFail(response http.ResponseWriter, msg string) {
	returnJson(response, -1, nil, msg)
}

func RespSuccess(response http.ResponseWriter, data interface{}, msg string) {
	returnJson(response, 0, data, msg)
}

// ReturnJson 返回json
func returnJson(response http.ResponseWriter, code int, data interface{}, msg string) {
	// header 为json
	response.Header().Set("Content-Type", "application/json")
	// 200状态码
	response.WriteHeader(http.StatusOK)

	result := ResponseMessage{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	ret, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
	}
	response.Write(ret)
}
