package services

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/model"
	"github.com/CanghaiLi/LearnPros/gin-pro/serializer"
)

type UserService struct {
	UserName string `json:"username" binding:"required, min=3,max=15"`
	Password string `json:"password" binding:"required, min=6,max=20"`
}

func (us *UserService) Register() serializer.Response {
	var user model.User
	var count int
	// 此时Register的调用，在shouldBind之后， 请求参数已经被绑定到UserService的实例 u 中了
	// 查user_name为请求参数username的一行数据，并绑到user变量上
	model.DB.Where("user_name=?", us.UserName).First(&user).Count(&count) // 是不是有一条这样的数据
	if count > 0 {
		// 证明用户名重复了
		return serializer.Response{
			Code: 400,
			Msg:  "用户名已存在",
		}
	}
	user.Username = us.UserName
	// 密码加密
	if err := user.EncryptPassword(us.Password); err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "密码错误" + err.Error(),
		}
	}
	// 创建用户。写库
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "数据库psd写入错误" + err.Error(),
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "OK",
	}

}
