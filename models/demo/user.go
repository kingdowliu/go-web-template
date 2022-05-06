package demo

import "gorm.io/gorm"

//
// User
// @Description: 用户结构体
//
type User struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}
