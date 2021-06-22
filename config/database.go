package config

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"hello/model"
	"log"
)

// xorm手册： https://gobook.io/read/gitea.com/xorm/manual-zh-CN/

type DbConfig struct {
	Type         string
	Hostname     string
	Port         int
	DatabaseName string
	UserName     string
	Password     string
	Character    string
}

const PAGE_SIZE = 50

var dbConfig *DbConfig

func init() {
	dbConfig = &DbConfig{
		Type:         "mysql",
		Hostname:     "127.0.0.1",
		Port:         3306,
		DatabaseName: "im",
		UserName:     "root",
		Password:     "123456",
		Character:    "utf8",
	}
}

func GetDbEngine() *xorm.Engine {
	//DatasourceName := "root:123456@(127.0.0.1:3306)/im?charset=utf8"
	dbsoure := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s",
		dbConfig.UserName, dbConfig.Password, dbConfig.Hostname, dbConfig.Port, dbConfig.DatabaseName, dbConfig.Character)
	err := errors.New("")
	var DbEngine *xorm.Engine
	DbEngine, err = xorm.NewEngine(dbConfig.Type, dbsoure)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	// 是否显示SQL语句
	DbEngine.ShowSQL(true)
	// 设置数据库最大连接数
	DbEngine.SetMaxOpenConns(2)
	// 自动User
	DbEngine.Sync2(new(model.User), new(model.Contact), new(model.Group), new(model.Community))
	fmt.Println("init database ok!")

	return DbEngine
}
