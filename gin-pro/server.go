package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"`
}

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
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
