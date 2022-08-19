package services

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/model"
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/response"
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/utils"
	"github.com/CanghaiLi/LearnPros/gin-pro/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (cts *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	// 找到数据库中
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     cts.Title,
		Content:   cts.Content,
		Status:    0, // 0完成， 1未完成
		StartTime: time.Now().Unix(),
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		utils.LoggersObj.Info(err)
		code := response.ErrorDatabase
		return serializer.Response{
			Code: code,
			Msg:  response.GetMsg(code),
		}
	}
	return serializer.SuccessResponse(serializer.BuildTask(task))
}

// ListTasksService 分页功能服务
type ListTasksService struct {
	Limit int `json:"limit"`
	Start int `json:"start"`
}

func (lts *ListTasksService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if lts.Limit == 0 {
		lts.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", id).Count(&total).
		Limit(lts.Limit).Offset((lts.Start - 1) * lts.Limit).
		Find(&tasks)
	return serializer.SuccessResponse(serializer.TasksInfo{
		Tasks: serializer.BuildTasks(tasks),
		Total: total,
	})
}
