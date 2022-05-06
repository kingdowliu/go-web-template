package demo

import (
	"github.com/kingdowliu/go-web-template/global"
	demoModel "github.com/kingdowliu/go-web-template/models/demo"
)

type UserDao struct {
}

func (userDao *UserDao) CreateUserDao(user demoModel.User) (err error) {
	err = global.DB.Create(&user).Error
	return err
}

func (userDao *UserDao) DeleteUserDao(user demoModel.User) (err error) {
	err = global.DB.Delete(&user).Error
	return err
}

func (userDao *UserDao) UpdateUserDao(user *demoModel.User) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func (userDao *UserDao) GetUserByIdDao(id uint) (user demoModel.User, err error) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return
}
