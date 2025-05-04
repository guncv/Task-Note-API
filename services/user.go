package services

import (
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/repositories"
)

type IUserService interface {
}

type UserService struct {
	repo repositories.IUserRepository
	log  *log.Logger
}

func NewUserService(repo repositories.IUserRepository, log *log.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}
