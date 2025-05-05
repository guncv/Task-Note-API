package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/models"
	"github.com/guncv/tech-exam-software-engineering/repositories"
	"github.com/guncv/tech-exam-software-engineering/utils"
	"gorm.io/gorm"
)

type IUserService interface {
	RegisterUser(ctx context.Context, req *entities.RegisterRequest) (*entities.RegisterResponse, error)
	LoginUser(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error)
}

type UserService struct {
	repo       repositories.IUserRepository
	log        *log.Logger
	tokenMaker utils.IPasetoMaker
	config     *config.Config
}

func NewUserService(
	repo repositories.IUserRepository,
	log *log.Logger,
	tokenMaker utils.IPasetoMaker,
	config *config.Config,
) IUserService {
	return &UserService{
		repo:       repo,
		log:        log,
		tokenMaker: tokenMaker,
		config:     config,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *entities.RegisterRequest) (*entities.RegisterResponse, error) {
	s.log.DebugWithID(ctx, "[Service: RegisterUser] Called")
	hashedPassword, err := utils.HashPassword(ctx, req.Password, s.log)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: RegisterUser] Failed to hash password: ", err)
		return nil, constants.ErrHashPassword
	}

	arg := models.User{
		ID:        uuid.New(),
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err = s.repo.RegisterUser(ctx, &arg); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			s.log.ErrorWithID(ctx, "[Service: RegisterUser] Duplicate user", err)
			return nil, constants.ErrUserAlreadyExists
		}

		s.log.ErrorWithID(ctx, "[Service: RegisterUser] Failed to register user: ", err)
		return nil, err
	}

	resp := entities.RegisterResponse{
		Id:           arg.ID.String(),
		FirstName:    arg.FirstName,
		LastName:     arg.LastName,
		Email:        arg.Email,
		HashPassword: hashedPassword,
	}

	s.log.DebugWithID(ctx, "[Service: RegisterUser] Successfully registered user: ", resp)
	return &resp, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error) {
	s.log.DebugWithID(ctx, "[Service: LoginUser] Called")

	user, err := s.repo.GetUser(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.ErrorWithID(ctx, "[Service: LoginUser] User not found: ", err)
			return nil, constants.ErrUserNotFound
		}
		s.log.ErrorWithID(ctx, "[Service: LoginUser] Failed to get user: ", err)
		return nil, err
	}

	if err = utils.CheckPassword(ctx, req.Password, user.Password, s.log); err != nil {
		s.log.ErrorWithID(ctx, "[Service: LoginUser] Failed to check password: ", err)
		return nil, constants.ErrPasswordIncorrect
	}

	accessToken, err := s.tokenMaker.CreateToken(
		user.ID.String(),
		s.config.TokenConfig.AccessTokenDuration,
	)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: LoginUser] Failed to create token: ", err)
		return nil, err
	}

	response := entities.LoginResponse{
		Token: accessToken,
	}

	s.log.DebugWithID(ctx, "[Service: LoginUser] Successfully logged in user: ", response)
	return &response, nil
}
