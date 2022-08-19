package api

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/utils"
	"github.com/CanghaiLi/LearnPros/gin-pro/services"
	"github.com/gin-gonic/gin"
)

func TaskCreate(ctx *gin.Context) {
	var taskCreate services.CreateTaskService
	chaim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&taskCreate); err == nil {
		res := taskCreate.Create(chaim.Id)
		ctx.JSON(200, res)
	} else {
		utils.LoggersObj.Info(err)
		ctx.JSON(400, err)
	}
}

func TasksQuery(ctx *gin.Context) {
	var listService services.ListTasksService
	chaim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&listService); err == nil {
		res := listService.List(chaim.Id)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, err)
		utils.LoggersObj.Info(err)
	}
}
