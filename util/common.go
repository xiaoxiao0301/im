package util

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func UploadImages(request *http.Request) (url string, err error) {
	srcFile, srcFileHeader, err := request.FormFile("file")
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	// 文件后缀
	var suffix string = ".png"
	// 查看上传的文件是佛包含文件后缀
	tmpfilename := srcFileHeader.Filename
	tmp := strings.Split(tmpfilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	// 查看前端是否传递了文件类型 formdata.append("filetype", "png")
	fileType := request.FormValue("filetype")
	if len(fileType) > 0 {
		suffix = fileType
	}

	// 构造上传后的文件名 时间戳.png
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), RandString(10), suffix)
	fmt.Println(filename)
	// 目标文件
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	path := fmt.Sprintf("./upload/%d%02d/%02d/", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	createMultiDir(path)
	//path := "./upload/"
	dstfile, err := os.OpenFile(path+filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer dstfile.Close()
	// 将上传的文件复制到目标地方
	_, err = io.Copy(dstfile, srcFile)
	if err != nil {
		return "", err
	}
	// 将新文件路径返回到前端
	fmt.Println(path)
	url = strings.Trim(path, ".") + filename
	return url, nil
}

// RandString 生成指定长度的随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// 创建文件夹
func createMultiDir(filePath string) error {
	if !isExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败", err.Error())
			return err
		}
		return err
	}
	return nil
}

// isExists 判断给定的文件或目录是否存在
func isExists(path string) bool {
	_, err := os.Stat(path) // 获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
