package services

import (
	"context"

	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/repositories"
)

type ITaskService interface {
	HealthCheck(ctx context.Context) (*entities.GetHealthUserResponse, error)
}

type TaskService struct {
	repo repositories.ITaskRepository
	log  *log.Logger
}

func NewTaskService(
	repo repositories.ITaskRepository,
	log *log.Logger,
) ITaskService {
	return &TaskService{
		repo: repo,
		log:  log,
	}
}

func (s *TaskService) HealthCheck(ctx context.Context) (*entities.GetHealthUserResponse, error) {
	s.log.DebugWithID(ctx, "[Service: HealthCheck] Called:")
	status, err := s.repo.HealthCheck(ctx)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: HealthCheck] Failed to health check")
		return nil, err
	}

	return &entities.GetHealthUserResponse{Status: status}, nil
}
