package routes

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/api"
	"github.com/CanghaiLi/LearnPros/gin-pro/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("lining-lsb"))
	r.Use(sessions.Sessions("mySession", store))
	v1 := r.Group("api/v1")
	{
		v1.GET("/test", api.Test)
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
		auth := v1.Group("/")
		auth.Use(middleware.JWT())
		{
			auth.POST("/task/create", api.TaskCreate)
			auth.POST("/task/list", api.TasksQuery)
		}
	}
	return r
}
