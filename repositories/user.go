package repositories

import (
	"context"

	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, email string) (*models.User, error)
}

type UserRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func NewUserRepository(db *gorm.DB, log *log.Logger) IUserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user *models.User) error {
	r.log.DebugWithID(ctx, "[Repository: RegisterUser] Called")
	err := r.db.Create(user).Error
	if err != nil {
		r.log.ErrorWithID(ctx, "[Repository: RegisterUser] Failed to register user", err)
		return err
	}

	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, email string) (*models.User, error) {
	r.log.DebugWithID(ctx, "[Repository: GetUser] Called")
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: GetUser] Failed to get user", err)
		return nil, err
	}

	return &user, nil
}
