package repositories

import (
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"gorm.io/gorm"
)

type IUserRepository interface {
}

type UserRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func NewUserRepository(db *gorm.DB, log *log.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}
