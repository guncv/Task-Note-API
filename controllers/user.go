package controllers

import (
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/services"
)

type IUserController interface {
}

type UserController struct {
	service services.IUserService
	log     *log.Logger
}

func NewUserController(service services.IUserService) *UserController {
	return &UserController{
		service: service,
	}
}
