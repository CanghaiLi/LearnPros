package middleware

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/utils"
	"github.com/CanghaiLi/LearnPros/gin-pro/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 9191
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 9191 // 错误的token 假token
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 9191 // 过期token
			}
		}
		if code != 200 {
			ctx.JSON(200, serializer.NewResponse(code, nil, "Token解析错误"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
