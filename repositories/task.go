package repositories

import (
	"context"

	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/models"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	HealthCheck(ctx context.Context) (string, error)
	CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
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

func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	r.log.DebugWithID(ctx, "[Repository: CreateTask] Called")

	if err := r.db.Create(task).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: CreateTask] Failed to create task", err)
		return nil, err
	}

	return task, nil
}
