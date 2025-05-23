package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/services"
	"github.com/guncv/tech-exam-software-engineering/utils"
)

type UserController struct {
	service services.IUserService
	log     *log.Logger
}

func NewUserController(service services.IUserService, log *log.Logger) *UserController {
	return &UserController{
		service: service,
		log:     log,
	}
}

// @Tags Users
// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param registerRequest body entities.RegisterRequest true "Register request"
// @Success 200 {object} entities.RegisterResponse "Successfully registered user"
// @Failure 400 {object} entities.ErrExampleInvalidRequest "Invalid request body"
// @Failure 409 {object} entities.ErrExampleUserExists "User already exists"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/users [post]
func (c *UserController) Register(ctx *gin.Context) {
	c.log.DebugWithID(ctx, "[Controller: Register] Called")
	var req entities.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.ErrorWithID(ctx, "[Controller: Register] Failed to bind request: ", err)
		detail := utils.ValidateRegisterInput(req)
		utils.ErrorResponse(ctx, constants.ErrInvalidRequestBody, detail)
		return
	}

	resp, err := c.service.RegisterUser(ctx, &req)
	if err != nil {
		c.log.ErrorWithID(ctx, "[Controller: Register] Failed to register user: ", err)
		utils.ErrorResponse(ctx, err)
		return
	}

	c.log.InfoWithID(ctx, "[Controller: Register] Successfully registered user")
	ctx.JSON(http.StatusOK, resp)
}

// @Tags Users
// @Summary Login a user
// @Description Login a user with email and password
// @Accept json
// @Produce json
// @Param loginRequest body entities.LoginRequest true "Login request"
// @Success 200 {object} entities.LoginResponse "Successfully logged in user"
// @Failure 400 {object} entities.ErrExampleInvalidRequest "Invalid request body"
// @Failure 401 {object} entities.ErrExampleIncorrectPassword "Password is incorrect"
// @Failure 404 {object} entities.ErrExampleUserNotFound "User not found"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	c.log.DebugWithID(ctx, "[Controller: Login] Called")
	var req entities.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := utils.ValidateLoginInput(req)
		c.log.ErrorWithID(ctx, "[Controller: Login] Failed to bind request: ", err)
		utils.ErrorResponse(ctx, constants.ErrInvalidRequestBody, detail)
		return
	}

	resp, err := c.service.LoginUser(ctx, &req)
	if err != nil {
		c.log.ErrorWithID(ctx, "[Controller: Login] Failed to login user: ", err)
		utils.ErrorResponse(ctx, err)
		return
	}

	c.log.InfoWithID(ctx, "[Controller: Login] Successfully logged in user")
	ctx.JSON(http.StatusOK, resp)
}
