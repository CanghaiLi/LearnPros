package services

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/model"
	"github.com/CanghaiLi/LearnPros/gin-pro/serializer"
)

type UserService struct {
	UserName string `json:"username" binding:"required, min=3,max=15"`
	Password string `json:"password" binding:"required, min=6,max=20"`
}

func (u *UserService) Register() serializer.Response {
	var user model.User
	var count int
	// 此时Register的调用，在shouldBind之后， 请求参数已经被绑定到UserService的实例 u 中了
	model.DB.Model(&model.User{}).
		Where("user_name=?", u.UserName).
		First(&user). // 查user_name为请求参数username的一行数据，并绑到user变量上
		Count(&count) // 是不是有一条这样的数据
	if count == 1 {
		// 证明用户名重复了
		return serializer.Response{
			Code: 400,
			Msg:  "用户名已存在",
		}
	}
	//user.Username =
	return serializer.Response{}
}
