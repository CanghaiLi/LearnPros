package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string // 密码，密文存储
}

func (u *User) EncryptPassword(ps string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ps), 12)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) ComparePassword(ps string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(ps))
	return err == nil
}
