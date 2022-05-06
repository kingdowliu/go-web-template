package demo

import (
	demoDao "github.com/kingdowliu/go-web-template/dao/demo"
	demoModel "github.com/kingdowliu/go-web-template/models/demo"
)

type UserService struct {
}

func (userService *UserService) CreateUserService(user demoModel.User) (err error) {
	var userDao demoDao.UserDao
	err = userDao.CreateUserDao(user)
	return err
}
