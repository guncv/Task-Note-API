package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/mocks"
	"github.com/guncv/tech-exam-software-engineering/models"
	"github.com/guncv/tech-exam-software-engineering/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskService_HealthCheck(t *testing.T) {
	errMockError := errors.New("mock error")
	lgr := log.Initialize(constants.TestAppEnv)

	ctx := context.Background()
	okResponse := "ok"
	okResponseEntity := &entities.GetHealthUserResponse{
		Status: okResponse,
	}

	testCases := []struct {
		name   string
		setup  func() *mocks.MockITaskRepository
		input  func() context.Context
		verify func(t *testing.T, got *entities.GetHealthUserResponse, gotErr error)
	}{
		{
			name: "HealthCheck_OK",
			setup: func() *mocks.MockITaskRepository {
				mockTaskRepo := new(mocks.MockITaskRepository)

				mockTaskRepo.EXPECT().
					HealthCheck(ctx).
					Return(okResponse, nil)
				return mockTaskRepo
			},
			input: func() context.Context {
				return ctx
			},
			verify: func(t *testing.T, got *entities.GetHealthUserResponse, gotErr error) {
				assert.Equal(t, okResponseEntity, got)
				assert.NoError(t, gotErr)
			},
		},
		{
			name: "HealthCheck_ReturnError",
			setup: func() *mocks.MockITaskRepository {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockTaskRepo.EXPECT().
					HealthCheck(ctx).
					Return("", errMockError)
				return mockTaskRepo
			},
			input: func() context.Context {
				return ctx
			},
			verify: func(t *testing.T, got *entities.GetHealthUserResponse, gotErr error) {
				assert.Equal(t, (*entities.GetHealthUserResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockTaskRepo := tC.setup()
			defer mockTaskRepo.AssertExpectations(t)

			svc := NewTaskService(mockTaskRepo, lgr, nil)

			got, gotErr := svc.HealthCheck(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}

func TestTaskService_CreateTask(t *testing.T) {
	lgr := log.Initialize(constants.TestAppEnv)
	ctx := context.Background()
	errMockError := errors.New("mock error")
	empty := ""
	okResponse := &entities.CreateTaskResponse{
		ID:          "1",
		UserID:      "1",
		Title:       "Test Task",
		Status:      "IN_PROGRESS",
		Description: nil,
		Image:       &empty,
		CreatedAt:   "2021-01-01",
	}

	okRequest := &entities.CreateTaskRequest{
		Title:       "Test Task",
		Description: nil,
		Image:       nil,
		Status:      "IN_PROGRESS",
	}

	okPayload := &utils.Payload{
		ID:        uuid.New(),
		UserId:    "1",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	testCases := []struct {
		name   string
		setup  func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct)
		input  func() (context.Context, *entities.CreateTaskRequest)
		verify func(t *testing.T, got *entities.CreateTaskResponse, gotErr error)
	}{
		{
			name: "CreateTask_OK",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					CreateTask(ctx, mock.MatchedBy(func(task *models.Task) bool {
						return task.Title == okRequest.Title &&
							task.Status == string(okRequest.Status) &&
							task.Description == okRequest.Description
					})).
					Return(nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, *entities.CreateTaskRequest) {
				return ctx, okRequest
			},
			verify: func(t *testing.T, got *entities.CreateTaskResponse, gotErr error) {
				assert.Equal(t, okResponse.Title, got.Title)
				assert.Equal(t, okResponse.Status, got.Status)
				assert.Equal(t, okResponse.Description, got.Description)
				assert.Equal(t, okResponse.Image, got.Image)
				assert.NoError(t, gotErr)
			},
		},
		{
			name: "CreateTask_AuthorizationError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, *entities.CreateTaskRequest) {
				return ctx, okRequest
			},
			verify: func(t *testing.T, got *entities.CreateTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.CreateTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "CreateTask_ReturnError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					CreateTask(ctx, mock.Anything).
					Return(errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, *entities.CreateTaskRequest) {
				return ctx, okRequest
			},
			verify: func(t *testing.T, got *entities.CreateTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.CreateTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockTaskRepo, mockPayload := tC.setup()
			defer mockTaskRepo.AssertExpectations(t)
			defer mockPayload.AssertExpectations(t)

			svc := NewTaskService(mockTaskRepo, lgr, mockPayload)

			got, gotErr := svc.CreateTask(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}

func TestTaskService_GetTask(t *testing.T) {
	lgr := log.Initialize(constants.TestAppEnv)
	ctx := context.Background()
	errMockError := errors.New("mock error")

	empty := ""
	requestId := "550e8400-e29b-41d4-a716-446655440000"

	okResponse := &entities.GetTaskResponse{
		ID:          requestId,
		UserID:      "1",
		Title:       "Test Task",
		Status:      "IN_PROGRESS",
		Description: nil,
		Image:       &empty,
		CreatedAt:   "2021-01-01",
	}

	okPayload := &utils.Payload{
		ID:        uuid.New(),
		UserId:    "1",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	testCases := []struct {
		name   string
		setup  func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct)
		input  func() (context.Context, string)
		verify func(t *testing.T, got *entities.GetTaskResponse, gotErr error)
	}{
		{
			name: "GetTask_OK",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				getTaskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      okPayload.UserId,
					Title:       "Test Task",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(getTaskResponse, nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, got *entities.GetTaskResponse, gotErr error) {
				assert.Equal(t, requestId, got.ID)
				assert.Equal(t, okPayload.UserId, got.UserID)
				assert.Equal(t, okResponse.Title, got.Title)
				assert.Equal(t, okResponse.Status, got.Status)
				assert.Equal(t, okResponse.Description, got.Description)
				assert.Equal(t, okResponse.Image, got.Image)
				assert.NoError(t, gotErr)
			},
		},
		{
			name: "GetTask_AuthorizationError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, got *entities.GetTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.GetTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "GetTask_ReturnError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, got *entities.GetTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.GetTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "GetTask_NotMatchUserID",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				getTaskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      "2",
					Title:       "Test Task",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(getTaskResponse, nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, got *entities.GetTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.GetTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockTaskRepo, mockPayload := tC.setup()
			defer mockTaskRepo.AssertExpectations(t)
			defer mockPayload.AssertExpectations(t)

			svc := NewTaskService(mockTaskRepo, lgr, mockPayload)

			got, gotErr := svc.GetTask(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}

func TestTaskService_UpdateTask(t *testing.T) {
	lgr := log.Initialize(constants.TestAppEnv)
	ctx := context.Background()
	errMockError := errors.New("mock error")

	empty := ""
	requestId := "550e8400-e29b-41d4-a716-446655440000"
	description := "Test Description"
	okResponse := &entities.UpdateTaskResponse{
		ID:          "1",
		UserID:      "1",
		Title:       "Test Tasks",
		Status:      "IN_PROGRESS",
		Description: &description,
		Image:       &empty,
		CreatedAt:   "2021-01-01",
	}

	okRequest := &entities.UpdateTaskRequest{
		Title:       "Test Tasks",
		Description: &description,
		Image:       nil,
		Status:      "IN_PROGRESS",
	}

	okPayload := &utils.Payload{
		ID:        uuid.New(),
		UserId:    "1",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	testCases := []struct {
		name   string
		setup  func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct)
		input  func() (context.Context, string, *entities.UpdateTaskRequest)
		verify func(t *testing.T, got *entities.UpdateTaskResponse, gotErr error)
	}{
		{
			name: "UpdateTask_OK",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				taskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      okPayload.UserId,
					Title:       "Test Tasks",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(taskResponse, nil)

				mockTaskRepo.EXPECT().
					UpdateTask(ctx, mock.Anything).
					Return(nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string, *entities.UpdateTaskRequest) {
				return ctx, requestId, okRequest
			},
			verify: func(t *testing.T, got *entities.UpdateTaskResponse, gotErr error) {
				assert.Equal(t, requestId, got.ID)
				assert.Equal(t, okPayload.UserId, got.UserID)
				assert.Equal(t, okResponse.Title, got.Title)
				assert.Equal(t, okResponse.Status, got.Status)
				assert.Equal(t, okResponse.Description, got.Description)
				assert.Equal(t, okResponse.Image, got.Image)
				assert.NoError(t, gotErr)
			},
		},
		{
			name: "UpdateTask_AuthorizationError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string, *entities.UpdateTaskRequest) {
				return ctx, requestId, okRequest
			},
			verify: func(t *testing.T, got *entities.UpdateTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.UpdateTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "UpdateTask_GetTaskError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string, *entities.UpdateTaskRequest) {
				return ctx, requestId, okRequest
			},
			verify: func(t *testing.T, got *entities.UpdateTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.UpdateTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "UpdateTask_NotMatchUserID",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				getTaskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      "2",
					Title:       "Test Task",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(getTaskResponse, nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string, *entities.UpdateTaskRequest) {
				return ctx, requestId, okRequest
			},
			verify: func(t *testing.T, got *entities.UpdateTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.UpdateTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "UpdateTask_UpdateTaskError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				taskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      okPayload.UserId,
					Title:       "Test Tasks",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(taskResponse, nil)

				mockTaskRepo.EXPECT().
					UpdateTask(ctx, mock.Anything).
					Return(errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string, *entities.UpdateTaskRequest) {
				return ctx, requestId, okRequest
			},
			verify: func(t *testing.T, got *entities.UpdateTaskResponse, gotErr error) {
				assert.Equal(t, (*entities.UpdateTaskResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockTaskRepo, mockPayload := tC.setup()
			defer mockTaskRepo.AssertExpectations(t)
			defer mockPayload.AssertExpectations(t)

			svc := NewTaskService(mockTaskRepo, lgr, mockPayload)

			got, gotErr := svc.UpdateTask(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}

func TestTaskService_DeleteTask(t *testing.T) {
	lgr := log.Initialize(constants.TestAppEnv)
	ctx := context.Background()
	errMockError := errors.New("mock error")
	empty := ""
	requestId := "550e8400-e29b-41d4-a716-446655440000"

	okPayload := &utils.Payload{
		ID:        uuid.New(),
		UserId:    "1",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	testCases := []struct {
		name   string
		setup  func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct)
		input  func() (context.Context, string)
		verify func(t *testing.T, gotErr error)
	}{
		{
			name: "DeleteTask_OK",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				taskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      okPayload.UserId,
					Title:       "Test Tasks",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(taskResponse, nil)

				mockTaskRepo.EXPECT().
					DeleteTask(ctx, mock.Anything).
					Return(nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, gotErr error) {
				assert.NoError(t, gotErr)
			},
		},
		{
			name: "DeleteTask_AuthorizationError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, gotErr error) {
				assert.Error(t, gotErr)
			},
		},
		{
			name: "DeleteTask_GetTaskError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, gotErr error) {
				assert.Error(t, gotErr)
			},
		},
		{
			name: "DeleteTask_NotMatchUserID",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				taskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      "2",
					Title:       "Test Tasks",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(taskResponse, nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, gotErr error) {
				assert.Error(t, gotErr)
			},
		},
		{
			name: "DeleteTask_DeleteTaskError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				taskResponse := &models.Task{
					ID:          uuid.MustParse(requestId),
					UserID:      okPayload.UserId,
					Title:       "Test Tasks",
					Status:      "IN_PROGRESS",
					Description: nil,
					Image:       &empty,
					CreatedAt:   time.Now(),
				}

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetTask(ctx, mock.MatchedBy(func(id string) bool {
						return id == requestId
					})).
					Return(taskResponse, nil)

				mockTaskRepo.EXPECT().
					DeleteTask(ctx, mock.Anything).
					Return(errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, string) {
				return ctx, requestId
			},
			verify: func(t *testing.T, gotErr error) {
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockTaskRepo, mockPayload := tC.setup()
			defer mockTaskRepo.AssertExpectations(t)
			defer mockPayload.AssertExpectations(t)

			svc := NewTaskService(mockTaskRepo, lgr, mockPayload)

			gotErr := svc.DeleteTask(tC.input())

			tC.verify(t, gotErr)
		})
	}
}

func TestTaskService_GetAllTasks(t *testing.T) {
	lgr := log.Initialize(constants.TestAppEnv)
	ctx := context.Background()
	errMockError := errors.New("mock error")

	okPayload := &utils.Payload{
		ID:        uuid.New(),
		UserId:    "1",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}

	testCases := []struct {
		name   string
		setup  func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct)
		input  func() (context.Context, *entities.GetAllTasksRequest)
		verify func(t *testing.T, got *entities.GetAllTasksResponse, gotErr error)
	}{
		{
			name: "GetAllTasks_OK",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTasks := &[]models.Task{
					{
						ID:          uuid.New(),
						UserID:      "1",
						Title:       "Task 1",
						Description: ptr("Description 1"),
						Image:       ptr("image1.png"),
						Status:      "pending",
						CreatedAt:   time.Now(),
					},
				}

				mockTaskRepo.EXPECT().
					GetAllTasks(ctx, mock.Anything, mock.Anything).
					Return(mockTasks, nil)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, *entities.GetAllTasksRequest) {
				return ctx, &entities.GetAllTasksRequest{}
			},
			verify: func(t *testing.T, got *entities.GetAllTasksResponse, gotErr error) {
				assert.NoError(t, gotErr)
				assert.NotNil(t, got)
				assert.Equal(t, 1, got.Total)
				assert.Equal(t, "Task 1", got.Tasks[0].Title)
			},
		},
		{
			name: "GetAllTasks_NoAuthorization",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, *entities.GetAllTasksRequest) {
				return ctx, &entities.GetAllTasksRequest{}
			},
			verify: func(t *testing.T, got *entities.GetAllTasksResponse, gotErr error) {
				assert.Error(t, gotErr)
				assert.Nil(t, got)
			},
		},
		{
			name: "GetAllTasks_GetAllTasksError",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayloadConstruct) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayloadConstruct)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					GetAllTasks(ctx, mock.Anything, mock.Anything).
					Return(nil, errMockError)

				return mockTaskRepo, mockPayload
			},
			input: func() (context.Context, *entities.GetAllTasksRequest) {
				return ctx, &entities.GetAllTasksRequest{}
			},
			verify: func(t *testing.T, got *entities.GetAllTasksResponse, gotErr error) {
				assert.Error(t, gotErr)
				assert.Nil(t, got)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockTaskRepo, mockPayload := tC.setup()
			defer mockTaskRepo.AssertExpectations(t)
			defer mockPayload.AssertExpectations(t)

			svc := NewTaskService(mockTaskRepo, lgr, mockPayload)

			got, gotErr := svc.GetAllTasks(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}

// Helper function to convert string to *string
func ptr(s string) *string {
	return &s
}
