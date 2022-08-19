package serializer

import (
	"github.com/CanghaiLi/LearnPros/gin-pro/model"
	"time"
)

type Task struct {
	ID        uint   `json:"id" example:"1"`       // 任务ID
	Title     string `json:"title" example:"吃饭"`   // 题目
	Content   string `json:"content" example:"睡觉"` // 内容
	View      uint64 `json:"view" example:"32"`    // 浏览量
	Status    int    `json:"status" example:"0"`   // 状态(0未完成，1已完成)
	CreatedAt string `json:"createdAt"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime,omitempty"`
}

type TasksInfo struct {
	Tasks []Task `json:"tasks"`
	Total int64  `json:"total"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
		Status:  item.Status,
		//View:      item.View(),
		CreatedAt: timeStamp2str(item.CreatedAt.Unix()),
		StartTime: timeStamp2str(item.StartTime),
	}
}

func timeStamp2str(stamp int64) string {
	return time.Unix(stamp, 0).Format("2006-01-02 15:04:05")
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
