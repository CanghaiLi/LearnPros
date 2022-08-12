package main

func main() {

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
//		result.Success([]interface{}{"周", "杰", "轮", 4, 5})
//	})
//	v1.POST("/test", func(ctx *gin.Context) {
//		// 注册检验函数
//		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//			v.RegisterValidation("range18", range18)
//		}
//		var u User
//		result := NewResult(ctx)
//		err := ctx.ShouldBind(&u)
//		if err != nil {
//			fmt.Println(err.Error())
//			result.Error(1, err.Error())
//		} else {
//			result.Success(&u)
//		}
//	})
//	r.Run()
//}
//
//func range18(fl validator.FieldLevel) bool {
//	if fl.Field().Interface().(int) >= 18 {
//		return true
//	}
//	return false
//}
