package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 对字符串进行md5加密
func Md5Encode(data string) string {
	secret := md5.New()
	secret.Write([]byte(data))
	cipherStr := secret.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

// MD5Encode 大写md5值
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// EncryptPwd 加密密码
func EncryptPwd(pwd, salt string) string {
	return MD5Encode(pwd + salt)
}

// CheckPwd 校验密码
func CheckPwd(pwd, salt, password string) bool {
	return MD5Encode(pwd+salt) == password
}
