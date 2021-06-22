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

type PageResponseMessage struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	CurrentPage int         `json:"current_page"`
	TotalPage   int         `json:"total_page"`
	Data        interface{} `json:"data"`
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

func RespSuccessList(response http.ResponseWriter, lists interface{}, totals int, page int, message string) {
	// 分页返回
	RespList(response, 0, lists, totals, page, message)
}

func RespList(response http.ResponseWriter, code int, data interface{}, total int, page int, message string) {
	// 设置返回格式为json
	response.Header().Set("Content-Type", "application/json")
	// 返回状态码为200
	response.WriteHeader(http.StatusOK)

	result := &PageResponseMessage{
		Code:        code,
		Message:     message,
		TotalPage:   total,
		CurrentPage: page,
		Data:        data,
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("分页输出出错:%s\n", err.Error())
	}
	response.Write(resultBytes)
}
