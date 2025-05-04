package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/mapping"
	"github.com/guncv/tech-exam-software-engineering/repositories"
)

type ITaskService interface {
	HealthCheck(ctx context.Context) (*entities.GetHealthUserResponse, error)
	CreateTask(ctx context.Context, req *entities.CreateTaskRequest) (*entities.CreateTaskResponse, error)
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

func (s *TaskService) CreateTask(ctx context.Context, req *entities.CreateTaskRequest) (*entities.CreateTaskResponse, error) {
	s.log.DebugWithID(ctx, "[Service: CreateTask] Called:")

	// Encode image to base64
	base64Image := ""
	if req.Image != nil {
		file, err := req.Image.Open()
		if err != nil {
			s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to open image", err)
			return nil, errors.New("unable to read image")
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, file); err != nil {
			s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to read image", err)
			return nil, errors.New("unable to read image")
		}

		base64Image = base64.StdEncoding.EncodeToString(buf.Bytes())
		s.log.DebugWithID(ctx, "[Service: CreateTask] Encoded image to Base64")
	}

	task := mapping.MapCreateTaskRequestToTask(req, base64Image, uuid.New())

	repoResponse, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to create task", err)
		return nil, errors.New("failed to create task")
	}

	resp := mapping.MapCreateTaskResponseToTask(repoResponse, "Task created successfully")

	s.log.DebugWithID(ctx, "[Service: CreateTask] Task created successfully", resp)
	return resp, nil
}
