package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/mocks"
	"github.com/guncv/tech-exam-software-engineering/models"
	"github.com/guncv/tech-exam-software-engineering/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_RegisterUser(t *testing.T) {
	errMockError := errors.New("mock error")
	lgr := log.Initialize(constants.TestAppEnv)

	ctx := context.Background()

	okResponseEntity := &entities.RegisterRequest{
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		Password:  "password_test",
	}

	testCases := []struct {
		name   string
		setup  func() *mocks.MockIUserRepository
		input  func() (context.Context, *entities.RegisterRequest)
		verify func(t *testing.T, got *entities.RegisterResponse, gotErr error)
	}{
		{
			name: "RegisterUser_OK",
			setup: func() *mocks.MockIUserRepository {
				mockUserRepo := new(mocks.MockIUserRepository)

				mockUserRepo.EXPECT().
					RegisterUser(ctx, mock.MatchedBy(func(user *models.User) bool {
						return user.Email == okResponseEntity.Email &&
							user.FirstName == okResponseEntity.FirstName &&
							user.LastName == okResponseEntity.LastName &&
							user.Password != ""
					})).
					Return(nil)
				return mockUserRepo
			},
			input: func() (context.Context, *entities.RegisterRequest) {
				return ctx, okResponseEntity
			},
			verify: func(t *testing.T, got *entities.RegisterResponse, gotErr error) {
				assert.Equal(t, okResponseEntity.Email, got.Email)
				assert.Equal(t, okResponseEntity.FirstName, got.FirstName)
				assert.Equal(t, okResponseEntity.LastName, got.LastName)

				assert.NoError(t, gotErr)
			},
		},
		{
			name: "RegisterUser_ReturnError",
			setup: func() *mocks.MockIUserRepository {
				mockUserRepo := new(mocks.MockIUserRepository)

				mockUserRepo.EXPECT().
					RegisterUser(ctx, mock.MatchedBy(func(user *models.User) bool {
						return user.Email == okResponseEntity.Email
					})).
					Return(errMockError)
				return mockUserRepo
			},

			input: func() (context.Context, *entities.RegisterRequest) {
				return ctx, okResponseEntity
			},
			verify: func(t *testing.T, got *entities.RegisterResponse, gotErr error) {
				assert.Equal(t, (*entities.RegisterResponse)(nil), got)
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockUserRepo := tC.setup()
			defer mockUserRepo.AssertExpectations(t)

			svc := NewUserService(mockUserRepo, lgr, nil, nil)

			got, gotErr := svc.RegisterUser(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}

func TestUserService_LoginUser(t *testing.T) {
	errMockError := errors.New("mock error")
	lgr := log.Initialize(constants.TestAppEnv)
	cfg := &config.Config{
		TokenConfig: config.TokenConfig{
			TokenSymmetricKey:   utils.RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}

	ctx := context.Background()
	loginRequestEntity := &entities.LoginRequest{
		Email:    "test@test.com",
		Password: "password_test",
	}

	fixedUserID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	testCases := []struct {
		name   string
		setup  func() (*mocks.MockIUserRepository, *mocks.MockIPasetoMaker)
		input  func() (context.Context, *entities.LoginRequest)
		verify func(t *testing.T, got *entities.LoginResponse, gotErr error)
	}{
		{
			name: "LoginUser_OK",
			setup: func() (*mocks.MockIUserRepository, *mocks.MockIPasetoMaker) {
				mockUserRepo := new(mocks.MockIUserRepository)
				mockTokenMaker := new(mocks.MockIPasetoMaker)

				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password_test"), bcrypt.DefaultCost)

				mockUserRepo.EXPECT().
					GetUser(ctx, loginRequestEntity.Email).
					Return(&models.User{
						ID:       fixedUserID,
						Email:    loginRequestEntity.Email,
						Password: string(hashedPassword),
					}, nil)

				mockTokenMaker.EXPECT().
					CreateToken(fixedUserID.String(), mock.AnythingOfType("time.Duration")).
					Return("mocked_token", nil)

				return mockUserRepo, mockTokenMaker
			},
			input: func() (context.Context, *entities.LoginRequest) {
				return ctx, loginRequestEntity
			},
			verify: func(t *testing.T, got *entities.LoginResponse, gotErr error) {
				assert.Equal(t, "mocked_token", got.Token)
				assert.NoError(t, gotErr)
			},
		},
		{
			name: "LoginUser_GetUserError",
			setup: func() (*mocks.MockIUserRepository, *mocks.MockIPasetoMaker) {
				mockUserRepo := new(mocks.MockIUserRepository)
				mockTokenMaker := new(mocks.MockIPasetoMaker)

				mockUserRepo.EXPECT().
					GetUser(ctx, loginRequestEntity.Email).
					Return(nil, errMockError)

				return mockUserRepo, mockTokenMaker
			},
			input: func() (context.Context, *entities.LoginRequest) {
				return ctx, loginRequestEntity
			},
			verify: func(t *testing.T, got *entities.LoginResponse, gotErr error) {
				assert.Nil(t, got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "LoginUser_CheckPasswordError",
			setup: func() (*mocks.MockIUserRepository, *mocks.MockIPasetoMaker) {
				mockUserRepo := new(mocks.MockIUserRepository)
				mockTokenMaker := new(mocks.MockIPasetoMaker)

				mockUserRepo.EXPECT().
					GetUser(ctx, loginRequestEntity.Email).
					Return(&models.User{
						ID:       fixedUserID,
						Email:    loginRequestEntity.Email,
						Password: "invalid_hash",
					}, nil)

				return mockUserRepo, mockTokenMaker
			},
			input: func() (context.Context, *entities.LoginRequest) {
				return ctx, loginRequestEntity
			},
			verify: func(t *testing.T, got *entities.LoginResponse, gotErr error) {
				assert.Nil(t, got)
				assert.Error(t, gotErr)
			},
		},
		{
			name: "LoginUser_CreateTokenError",
			setup: func() (*mocks.MockIUserRepository, *mocks.MockIPasetoMaker) {
				mockUserRepo := new(mocks.MockIUserRepository)
				mockTokenMaker := new(mocks.MockIPasetoMaker)

				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password_test"), bcrypt.DefaultCost)

				mockUserRepo.EXPECT().
					GetUser(ctx, loginRequestEntity.Email).
					Return(&models.User{
						ID:       fixedUserID,
						Email:    loginRequestEntity.Email,
						Password: string(hashedPassword),
					}, nil)

				mockTokenMaker.EXPECT().
					CreateToken(fixedUserID.String(), mock.AnythingOfType("time.Duration")).
					Return("", errMockError)

				return mockUserRepo, mockTokenMaker
			},
			input: func() (context.Context, *entities.LoginRequest) {
				return ctx, loginRequestEntity
			},
			verify: func(t *testing.T, got *entities.LoginResponse, gotErr error) {
				assert.Nil(t, got)
				assert.Error(t, gotErr)
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mockUserRepo, mockTokenMaker := tC.setup()
			defer mockUserRepo.AssertExpectations(t)
			defer mockTokenMaker.AssertExpectations(t)

			svc := NewUserService(mockUserRepo, lgr, mockTokenMaker, cfg)

			got, gotErr := svc.LoginUser(tC.input())

			tC.verify(t, got, gotErr)
		})
	}
}
