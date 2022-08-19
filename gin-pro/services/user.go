package services

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/model"
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/response"
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/utils"
	"github.com/CanghaiLi/LearnPros/gin-pro/serializer"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `json:"username" binding:"required,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

func (us *UserService) Register() serializer.Response {
	code := response.SUCCESS
	var user model.User
	var count int64
	// 此时Register的调用，在shouldBind之后， 请求参数已经被绑定到UserService的实例 u 中了
	// 查user_name为请求参数username的一行数据，并绑到user变量上
	model.DB.Where("user_name=?", us.UserName).First(&user).Count(&count) // 是不是有一条这样的数据
	if count > 0 {
		// 证明用户名重复了
		code = response.ErrorExistUser
		return serializer.Response{
			Code: code,
			Msg:  "用户名已存在",
		}
	}
	user.UserName = us.UserName
	// 密码加密
	if err := user.EncryptPassword(us.Password); err != nil {
		utils.LoggersObj.Info(err)
		code = response.ErrorFailEncryption
		return serializer.Response{
			Code: code,
			Msg:  "密码格式错误，无法加密" + err.Error(),
		}
	}
	// 创建用户，写库
	if err := model.DB.Create(&user).Error; err != nil {
		utils.LoggersObj.Info(err)
		code = response.ErrorDatabase
		return serializer.Response{
			Code: code,
			Msg:  "数据库psd写入错误" + err.Error(),
		}
	}
	return serializer.Response{
		Code: code,
		Msg:  "OK",
	}
}

func (us *UserService) Login() serializer.Response {

	var user model.User
	// 先查询数据库是否存在这个人 username
	if err := model.DB.Where("user_name=?", us.UserName).First(&user).Error; err != nil {
		// 如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			utils.LoggersObj.Info(err)
			code := response.ErrorNotExistUser
			return serializer.Response{
				Code: code,
				Msg:  response.GetMsg(code),
			}
		}
		// 查询出现的另外错误
		utils.LoggersObj.Info(err)
		code := response.ErrorDatabase
		return serializer.Response{
			Code: code,
			Msg:  response.GetMsg(code),
		}
	}

	// 验证密码
	if !user.ComparePassword(us.Password) {
		code := response.ErrorNotCompare
		return serializer.Response{
			Code: code,
			Msg:  response.GetMsg(code),
		}
	}
	// 验证成功，下发token
	token, err := utils.GenerateToken(user.ID, user.UserName, us.Password)
	if err != nil {
		code := response.ErrorAuthToken
		return serializer.Response{
			Code: code,
			Msg:  response.GetMsg(code),
		}
	}
	return serializer.SuccessResponse(token)
}
