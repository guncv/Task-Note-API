package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/models"
	"github.com/guncv/tech-exam-software-engineering/repositories"
	"github.com/guncv/tech-exam-software-engineering/utils"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ITaskService interface {
	HealthCheck(ctx context.Context) (*entities.GetHealthUserResponse, error)
	CreateTask(ctx context.Context, req *entities.CreateTaskRequest) (*entities.CreateTaskResponse, error)
	GetTask(ctx context.Context, req string) (*entities.GetTaskResponse, error)
	UpdateTask(ctx context.Context, id string, req *entities.UpdateTaskRequest) (*entities.UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, id string) error
	GetAllTasks(ctx context.Context, req *entities.GetAllTasksRequest) (*entities.GetAllTasksResponse, error)
}

type TaskService struct {
	repo    repositories.ITaskRepository
	log     *log.Logger
	payload utils.IPayloadConstruct
}

func NewTaskService(
	repo repositories.ITaskRepository,
	log *log.Logger,
	payload utils.IPayloadConstruct,
) ITaskService {
	return &TaskService{
		repo:    repo,
		log:     log,
		payload: payload,
	}
}

func (s *TaskService) HealthCheck(ctx context.Context) (*entities.GetHealthUserResponse, error) {
	s.log.DebugWithID(ctx, "[Service: HealthCheck] Called:")
	status, err := s.repo.HealthCheck(ctx)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: HealthCheck] Failed to health check", err)
		return nil, err
	}

	return &entities.GetHealthUserResponse{Status: status}, nil
}

func (s *TaskService) CreateTask(ctx context.Context, req *entities.CreateTaskRequest) (*entities.CreateTaskResponse, error) {
	s.log.DebugWithID(ctx, "[Service: CreateTask] Called:")

	// Get auth payload
	authPayload, err := s.payload.GetAuthPayload(ctx, s.log)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to get auth payload", err)
		return nil, err
	}
	s.log.DebugWithID(ctx, "[Service: CreateTask] Auth payload: ", authPayload)

	// Encode image to base64
	base64Image := ""
	if req.Image != nil {
		var err error
		base64Image, err = utils.ConvertFileHeaderToBase64(req.Image)
		if err != nil {
			s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to encode image to base64", err)
			return nil, err
		}
	}

	// Create task
	arg := &models.Task{
		ID:          uuid.New(),
		UserID:      authPayload.UserId,
		Title:       req.Title,
		Status:      string(req.Status),
		Image:       &base64Image,
		Date:        req.Date,
		Description: req.Description,
		CreatedAt:   utils.NowBangkok(),
	}
	s.log.DebugWithID(ctx, "[Service: CreateTask] Task: ", arg)

	// Create task in repository
	if err := s.repo.CreateTask(ctx, arg); err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to create task", err)
				return nil, constants.ErrTaskAlreadyExists
			}
		}
		s.log.ErrorWithID(ctx, "[Service: CreateTask] Failed to create task", err)
		return nil, err
	}

	// Convert to response
	resp := &entities.CreateTaskResponse{
		ID:          arg.ID.String(),
		UserID:      arg.UserID,
		Title:       arg.Title,
		Status:      arg.Status,
		Date:        arg.Date,
		Image:       arg.Image,
		Description: arg.Description,
		CreatedAt:   arg.CreatedAt,
	}

	s.log.DebugWithID(ctx, "[Service: CreateTask] Task created successfully", resp)
	return resp, nil
}

func (s *TaskService) GetTask(ctx context.Context, req string) (*entities.GetTaskResponse, error) {
	s.log.DebugWithID(ctx, "[Service: GetTask] Called:")

	// Get auth payload
	authPayload, err := s.payload.GetAuthPayload(ctx, s.log)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetTask] Failed to get auth payload", err)
		return nil, err
	}
	s.log.DebugWithID(ctx, "[Service: GetTask] Auth payload: ", authPayload)

	// Get task from repository
	repoResponse, err := s.repo.GetTask(ctx, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.ErrorWithID(ctx, "[Service: GetTask] Task not found: ", err)
			return nil, constants.ErrTaskNotFound
		}

		s.log.ErrorWithID(ctx, "[Service: GetTask] Failed to get task", err)
		return nil, err
	}

	// Verify payload user id with your account id
	if authPayload.UserId != repoResponse.UserID {
		s.log.ErrorWithID(ctx, "[Service: GetTask] Failed to verify task id", constants.ErrUserIdDoesNotMatchWithYourAccount)
		return nil, constants.ErrUserIdDoesNotMatchWithYourAccount
	}

	// Convert to response
	resp := &entities.GetTaskResponse{
		ID:          repoResponse.ID.String(),
		UserID:      repoResponse.UserID,
		Title:       repoResponse.Title,
		Status:      repoResponse.Status,
		Image:       repoResponse.Image,
		Date:        repoResponse.Date,
		Description: repoResponse.Description,
		CreatedAt:   repoResponse.CreatedAt,
	}

	s.log.DebugWithID(ctx, "[Service: GetTask] Task retrieved successfully", resp)
	return resp, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id string, req *entities.UpdateTaskRequest) (*entities.UpdateTaskResponse, error) {
	s.log.DebugWithID(ctx, "[Service: UpdateTask] Called")

	// Get auth payload
	authPayload, err := s.payload.GetAuthPayload(ctx, s.log)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: UpdateTask] Failed to get auth payload", err)
		return nil, err
	}

	// Get existing task
	existingTask, err := s.repo.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.ErrorWithID(ctx, "[Service: UpdateTask] Task not found: ", err)
			return nil, constants.ErrTaskNotFound
		}

		s.log.ErrorWithID(ctx, "[Service: UpdateTask] Failed to get task", err)
		return nil, err
	}

	// Check ownership
	if authPayload.UserId != existingTask.UserID {
		s.log.ErrorWithID(ctx, "[Service: UpdateTask] User ID does not match task owner", constants.ErrUserIdDoesNotMatchWithYourAccount)
		return nil, constants.ErrUserIdDoesNotMatchWithYourAccount
	}

	// Update fields if present
	if req.Title != "" {
		existingTask.Title = req.Title
	}
	if req.Description != nil {
		existingTask.Description = req.Description
	}
	if req.Status != "" {
		existingTask.Status = string(req.Status)
	}
	if !req.Date.IsZero() {
		existingTask.Date = req.Date
	}
	if req.Image != nil {
		base64Image, err := utils.ConvertFileHeaderToBase64(req.Image)
		if err != nil {
			s.log.ErrorWithID(ctx, "[Service: UpdateTask] Failed to convert image to base64", err)
			return nil, err
		}
		existingTask.Image = &base64Image
	}

	// Save update
	if err := s.repo.UpdateTask(ctx, existingTask); err != nil {
		s.log.ErrorWithID(ctx, "[Service: UpdateTask] Failed to update task", err)
		return nil, err
	}

	// Response
	resp := &entities.UpdateTaskResponse{
		ID:          existingTask.ID.String(),
		UserID:      existingTask.UserID,
		Title:       existingTask.Title,
		Status:      existingTask.Status,
		Date:        existingTask.Date,
		Image:       existingTask.Image,
		Description: existingTask.Description,
		CreatedAt:   existingTask.CreatedAt,
	}

	s.log.DebugWithID(ctx, "[Service: UpdateTask] Task updated successfully", resp)
	return resp, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	s.log.DebugWithID(ctx, "[Service: DeleteTask] Called")

	// Get auth payload
	authPayload, err := s.payload.GetAuthPayload(ctx, s.log)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: DeleteTask] Failed to get auth payload", err)
		return err
	}

	// Get existing task
	existingTask, err := s.repo.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.ErrorWithID(ctx, "[Service: DeleteTask] Task not found: ", err)
			return constants.ErrTaskNotFound
		}

		s.log.ErrorWithID(ctx, "[Service: DeleteTask] Failed to get task", err)
		return err
	}

	// Check ownership
	if authPayload.UserId != existingTask.UserID {
		s.log.ErrorWithID(ctx, "[Service: DeleteTask] User ID does not match task owner", constants.ErrUserIdDoesNotMatchWithYourAccount)
		return constants.ErrUserIdDoesNotMatchWithYourAccount
	}

	// Delete task
	if err := s.repo.DeleteTask(ctx, id); err != nil {
		s.log.ErrorWithID(ctx, "[Service: DeleteTask] Failed to delete task", err)
		return err
	}

	s.log.DebugWithID(ctx, "[Service: DeleteTask] Task deleted successfully")
	return nil
}

func (s *TaskService) GetAllTasks(ctx context.Context, req *entities.GetAllTasksRequest) (*entities.GetAllTasksResponse, error) {
	s.log.DebugWithID(ctx, "[Service: GetAllTasks] Called")

	// Get auth payload
	authPayload, err := s.payload.GetAuthPayload(ctx, s.log)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetAllTasks] Failed to get auth payload", err)
		return nil, err
	}
	s.log.DebugWithID(ctx, "[Service: GetAllTasks] Auth payload: ", authPayload)

	// Apply default values
	if req.Order == "" {
		req.Order = "desc"
	}
	if req.SortBy == "" {
		req.SortBy = "created_at"
	}

	// Set Offset
	req.Offset = (req.Offset - 1) * req.Limit

	// Fetch tasks from repository
	repoTasks, err := s.repo.GetAllTasks(ctx, req, authPayload.UserId)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetAllTasks] Failed to get all tasks", err)
		return nil, err
	}

	// Convert []models.Task → []entities.Task
	var tasks []entities.GetTaskResponse
	for _, t := range *repoTasks {
		tasks = append(tasks, entities.GetTaskResponse{
			ID:          t.ID.String(),
			UserID:      t.UserID,
			Title:       t.Title,
			Description: t.Description,
			Image:       t.Image,
			Date:        t.Date,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt,
		})
	}

	response := &entities.GetAllTasksResponse{
		Total: len(tasks),
		Tasks: tasks,
	}

	s.log.DebugWithID(ctx, "[Service: GetAllTasks] Tasks retrieved successfully", response)
	return response, nil
}
