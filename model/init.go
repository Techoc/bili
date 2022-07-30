package model

import (
	"github.com/techoc/bili/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

func Init() {
	util.Log().Info("初始化数据库连接")

	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	//db.SetLogger(util.Log())
	if err != nil {
		util.Log().Panic("连接数据库不成功, %s", err)
	}

	//设置连接池
	sqlite, err := db.DB()
	sqlite.SetMaxIdleConns(50)
	sqlite.SetMaxOpenConns(1)

	//超时
	sqlite.SetConnMaxLifetime(time.Second * 30)

	DB = db

	//执行迁移
	DB.AutoMigrate(&User{})
}
