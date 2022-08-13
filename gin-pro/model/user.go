package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	PasswordDigest string // 密码，密文存储
}
