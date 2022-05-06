package demo

import (
	demoModel "github.com/kingdowliu/go-web-template/models/demo"
	demoService "github.com/kingdowliu/go-web-template/service/demo"
)

type UserApi struct{}

func (userApi *UserApi) CreateUser() {
	user := demoModel.User{
		UserName: "test",
		PassWord: "test123",
	}
	var userService demoService.UserService
	err := userService.CreateUserService(user)
	if err != nil {
		return
	}
}
