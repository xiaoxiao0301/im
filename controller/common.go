package controller

import (
	"hello/util"
	"net/http"
)

// UploadFileHelper 上传文件
func UploadFileHelper(response http.ResponseWriter, request *http.Request) {
	url, err := util.UploadImages(request)
	if err != nil {
		util.RespFail(response, err.Error())
	} else {
		util.RespSuccess(response, url, "文件上传成功!")
	}
}
