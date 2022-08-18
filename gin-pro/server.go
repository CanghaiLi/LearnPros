package main

import (
	"fmt"
	"github.com/CanghaiLi/LearnPros/gin-pro/conf"
	"github.com/CanghaiLi/LearnPros/gin-pro/routes"
	"os"
)

func main() {
	fmt.Println(`os.GetEnv("JWT_SECRET")`, os.Getenv("JWT_SECRET"))
	conf.Init()
	r := routes.NewRouter()
	r.Run(":3000")
}

//type User struct {
//	Name string `json:"name"`
//	Age  int    `json:"age" binding:"required,range18"`
//	Sex  bool   `json:"sex"`
//}
//
//func main() {
//	f, _ := os.Create("gin.log")
//	gin.DefaultWriter = io.MultiWriter(f)
//	// 创建基础路由器
//	r := gin.Default()
//	v1 := r.Group("api/pc")
//	v1.GET("/keys", func(ctx *gin.Context) {
//		// 获取前端的header参数
//		token := ctx.Request.Header.Get("x-token")
//		fmt.Println(token)
//		result := NewResult(ctx)
//		var user User
//		user.Name = "lee"
//		user.Age = 18
//		user.Sex = true
//		result.Success(&user)
//	})
//	v1.POST("/test", func(ctx *gin.Context) {
//		// 注册检验函数
//		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//			v.RegisterValidation("range18", range18)
//		}
//		var u User
//		result := NewResult(ctx)
//		response := ctx.ShouldBind(&u)
//		if response != nil {
//			fmt.Println(response.Error())
//			result.Error(1, response.Error())
//		} else {
//			result.Success(&u)
//		}
//	})
//	r.Run()
//}

//func range18(fl validator.FieldLevel) bool {
//	if fl.Field().Interface().(int) >= 18 {
//		return true
//	}
//	return false
//}
