package repositories

import (
	"context"

	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	HealthCheck(ctx context.Context) (string, error)
}

type TaskRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func NewTaskRepository(
	db *gorm.DB,
	log *log.Logger,
) ITaskRepository {
	return &TaskRepository{
		db:  db,
		log: log,
	}
}

func (r *TaskRepository) HealthCheck(ctx context.Context) (string, error) {
	r.log.DebugWithID(ctx, "[Repository: HealthCheck] Called")
	return "Healthy", nil
}
