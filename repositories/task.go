package repositories

import (
	"context"

	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/models"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	HealthCheck(ctx context.Context) (string, error)
	CreateTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, id string) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id string) error
	GetAllTasks(ctx context.Context, req *entities.GetAllTasksRequest, userId string) (*[]models.Task, error)
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

func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	r.log.DebugWithID(ctx, "[Repository: CreateTask] Called")

	if err := r.db.Create(task).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: CreateTask] Failed to create task", err)
		return err
	}

	return nil
}

func (r *TaskRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	r.log.DebugWithID(ctx, "[Repository: GetTask] Called")

	var task models.Task
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: GetTask] Failed to get task", err)
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	r.log.DebugWithID(ctx, "[Repository: UpdateTask] Called")

	if err := r.db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(task).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: UpdateTask] Failed to update task", err)
		return err
	}

	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id string) error {
	r.log.DebugWithID(ctx, "[Repository: DeleteTask] Called")

	if err := r.db.Where("id = ?", id).Delete(&models.Task{}).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: DeleteTask] Failed to delete task", err)
		return err
	}

	return nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context, req *entities.GetAllTasksRequest, userId string) (*[]models.Task, error) {
	r.log.DebugWithID(ctx, "[Repository: GetAllTasks] Called")

	var tasks []models.Task
	query := r.db.Where("(title LIKE ? OR description LIKE ?) AND user_id = ?", "%"+req.Search+"%", "%"+req.Search+"%", userId)

	if err := query.Order(req.SortBy + " " + req.Order).Limit(req.Limit).Offset(req.Offset).Find(&tasks).Error; err != nil {
		r.log.ErrorWithID(ctx, "[Repository: GetAllTasks] Failed to get all tasks", err)
		return nil, err
	}

	return &tasks, nil
}
