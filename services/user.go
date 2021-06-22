package services

import (
	"errors"
	"fmt"
	"hello/config"
	"hello/model"
	"hello/util"
	"math/rand"
	"time"
)

type UserService struct {
}

// RegisterUser 注册
func (this *UserService) RegisterUser(
	mobile, //手机号
	password, // 明文密码
	nickname, // 昵称
	avatar, // 头像
	sex string) (user model.User, err error) {
	// 检测手机号是否存在
	user = model.User{}
	_, err = config.GetDbEngine().Where("mobile=?", mobile).Get(&user)
	if err != nil {
		return user, err
	}
	if user.Id > 0 {
		return user, errors.New("该手机号已注册")
	}
	user.Mobile = mobile
	user.Avatar = avatar
	user.Nickname = nickname
	user.Sex = sex
	salt := fmt.Sprintf("%06d", rand.Int31n(10000))
	user.Salt = salt
	user.Password = util.EncryptPwd(password, salt)
	user.CreateAt = time.Now()
	user.Token = fmt.Sprintf("%08d", rand.Int31())

	// 前端恶意插入字符，数据库操作失败
	_, insertErr := config.GetDbEngine().InsertOne(&user)
	return user, insertErr
}

// LoginUser 登录
func (this *UserService) LoginUser(
	mobile, //手机号
	password string) (user model.User, err error) {

	user = model.User{}
	oks, err := config.GetDbEngine().Where("mobile=?", mobile).Get(&user)

	if err != nil {
		return user, err
	}

	if !oks {
		return user, errors.New("手机号不存在!")
	}
	// 判断密码是否正确
	flag := util.CheckPwd(password, user.Salt, user.Password)
	if flag {
		// 刷新token
		str := fmt.Sprintf("%d", time.Now().Unix())
		token := util.MD5Encode(str)
		user.Token = token
		config.GetDbEngine().Id(user.Id).Cols("token").Update(&user)
		return user, nil
	} else {
		return user, errors.New("手机号或密码错误")
	}
}

// FindUser 查找用户
func (this *UserService) FindUser(userid int64) (user model.User) {
	config.GetDbEngine().Where("id = ?", userid).Get(&user)
	return user
}
