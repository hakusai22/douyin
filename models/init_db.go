package models

import (
	"github.com/hakusai22/douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局参数 DB
var DB *gorm.DB

// InitDB 初始化DB
func InitDB() {
	var err error
	// gorm 进行mysql连接 返回DB+error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}
	//自动根据结构体创建表 判断是否返回error
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
}
