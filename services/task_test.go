package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/mocks"
	"github.com/guncv/tech-exam-software-engineering/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskService_HealthCheck(t *testing.T) {
	errMockError := errors.New("mock error")
	lgr := log.Initialize("local")

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
	lgr := log.Initialize("local")

	ctx := context.Background()

	okResponse := &entities.CreateTaskResponse{
		ID:          "1",
		UserID:      "1",
		Title:       "Test Task",
		Status:      "IN_PROGRESS",
		Description: nil,
		Image:       nil,
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
		setup  func() (*mocks.MockITaskRepository, *mocks.MockIPayload)
		input  func() (context.Context, *entities.CreateTaskRequest)
		verify func(t *testing.T, got *entities.CreateTaskResponse, gotErr error)
	}{
		{
			name: "CreateTask_OK",
			setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayload) {
				mockTaskRepo := new(mocks.MockITaskRepository)
				mockPayload := new(mocks.MockIPayload)

				mockPayload.EXPECT().
					GetAuthPayload(ctx, mock.Anything).
					Return(okPayload, nil)

				mockTaskRepo.EXPECT().
					CreateTask(ctx, mock.Anything).
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
		// {
		// 	name: "CreateTask_ReturnError",
		// 	setup: func() (*mocks.MockITaskRepository, *mocks.MockIPayload) {
		// 		mockTaskRepo := new(mocks.MockITaskRepository)
		// 		mockPayload := new(mocks.MockIPayload)

		// 		mockTaskRepo.EXPECT().
		// 			CreateTask(ctx, mock.Anything).
		// 			Return("", errMockError)
		// 		return mockTaskRepo, mockPayload
		// 	},
		// 	input: func() (context.Context, *entities.CreateTaskRequest) {
		// 		return ctx, &entities.CreateTaskRequest{
		// 			Title:       "Test Task",
		// 			Description: "Test Description",
		// 			Image:       "Test Image",
		// 		}
		// 	},
		// 	verify: func(t *testing.T, got *entities.CreateTaskResponse, gotErr error) {
		// 		assert.Equal(t, (*entities.CreateTaskResponse)(nil), got)
		// 		assert.Error(t, gotErr)
		// 	},
		// },
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
