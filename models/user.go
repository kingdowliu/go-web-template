package models

import "gorm.io/gorm"

//
// User
// @Description: 用户结构体
//
type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"comment:用户名"`
	PassWord string `json:"password" gorm:"密码"`
}
