package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"required,range18"`
	Sex  bool   `json:"sex"`
}

func ccc[T int](p T) T {
	return p
}

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {

		// 注册检验函数
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("range18", range18)
		}

		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("%s", err),
			})
		} else {
			c.JSON(http.StatusOK, u)
		}
	})
	r.Run()
}

func range18(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(int) >= 18 {
		return true
	}
	return false

}
