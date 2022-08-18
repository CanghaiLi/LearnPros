package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"foreignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:'0'"` // 0未完成， 1已完成
	Content   string `gorm:"type:longtext"`
	StartTime int64  // 任务开始时间
	EndTime   int64  // 任务完成时间
}
