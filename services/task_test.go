package services

import (
	"context"
	"errors"
	"testing"

	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/mocks"
	"github.com/stretchr/testify/assert"
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

			svc := NewTaskService(mockTaskRepo, lgr)

			got, gotErr := svc.HealthCheck(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}
