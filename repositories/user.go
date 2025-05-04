package repositories

import (
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(user *models.User) error
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

func (r *UserRepository) RegisterUser(user *models.User) error {
	r.log.Info("Registering user", "user", user)
	err := r.db.Create(user).Error
	if err != nil {
		r.log.Error("Failed to register user", "error", err)
		return err
	}
	return nil
}
