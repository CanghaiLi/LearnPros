package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/big"
	"strconv"
	"time"
)

type Student struct {
	gorm.Model
	name string
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func Time2String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	db, _ := gorm.Open("mysql", "root:ln941001@/user_db?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	//自动迁移，也就是初始化数据表
	//db.AutoMigrate(&User{})
	//
	user := User{Name: "lee", Age: 28}
	user.CreateTime = Time2String(time.Now())
	user.UpdateTime = Time2String(time.Now())
	////db.NewRecord(user) // 主键为空返回`true`
	//db.Create(&user) // 创建user
	////db.NewRecord(user) // 创建`user`后返回`false`
	//
	var us []User
	var count int
	db.Debug().Where("create_time > ?", "2022-08-15 11:00:00").Find(&us).Count(&count)
	fmt.Println(count)
	//marshal, _ := json.Marshal(us)
	//fmt.Println(string(marshal))
	for i, b := range us {
		stream, _ := json.Marshal(b)
		fmt.Println(i, string(stream))
	}
	TestListComments()
}

func TestListComments() {
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	from, _ := time.ParseInLocation(layout, "2022-08-15 11:18:15", loc) //不加载时区
	fromUnix := from.Unix()
	fmt.Println(from)
	fmt.Println(fromUnix)
	to := time.Now()
	toUnix := to.Unix()
	fmt.Println(to)
	fmt.Println(toUnix)
	var v = 1.232384848484812841823123
	v2 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Println(v2)
	x, y := big.NewRat(1, 2), big.NewRat(2, 3)
	fmt.Println(x, y)
	z := new(big.Rat).Mul(x, y)
	fmt.Println("z:", z)
	fmt.Println("com:", (1/2)*(2/3))
}
