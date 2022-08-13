package api

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/services"
	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userRegister services.UserService
	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, err)
	}
}
