package routes

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("lining-lsb"))
	r.Use(sessions.Sessions("mySession", store))
	v1 := r.Group("api/pc")
	{
		v1.POST("user/register", api.UserRegister)
	}
	return r
}
