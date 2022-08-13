package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB

func DBConnect(dsn string) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库链接出错", err)
	}
	fmt.Println("数据库连接成功啦0=0")
	// 数据库日志输出 true详细版
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	// 别给表名自动加s
	db.SingularTable(true)
	// 设置连接数
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(40)
	db.DB().SetConnMaxLifetime(time.Second * 10)
	DB = db
	migration()
}
