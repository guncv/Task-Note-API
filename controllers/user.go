package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/services"
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

// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param registerRequest body entities.RegisterRequest true "Register request"
// @Success 200 {object} entities.RegisterResponse
// @Failure 400 {object} gin.H{"error": "Invalid request"}
// @Failure 500 {object} gin.H{"error": "Failed to register user"}
// @Router /v1/users/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	c.log.DebugWithID(ctx, "[Controller: Register] Called")
	var req entities.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.ErrorWithID(ctx, "[Controller: Register] Failed to bind request: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := c.service.RegisterUser(ctx, &req)
	if err != nil {
		c.log.ErrorWithID(ctx, "[Controller: Register] Failed to register user: ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	c.log.InfoWithID(ctx, "[Controller: Register] Successfully registered user")
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Login a user
// @Description Login a user with email and password
// @Accept json
// @Produce json
// @Param loginRequest body entities.LoginRequest true "Login request"
// @Success 200 {object} entities.LoginResponse
// @Failure 400 {object} gin.H{"error": "Invalid request"}
// @Failure 401 {object} gin.H{"error": "Invalid credentials"}
// @Failure 500 {object} gin.H{"error": "Failed to login user"}
// @Router /v1/users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	c.log.DebugWithID(ctx, "[Controller: Login] Called")
	var req entities.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.ErrorWithID(ctx, "[Controller: Login] Failed to bind request: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.service.LoginUser(ctx, &req)
	if err != nil {
		if err == sql.ErrNoRows {
			c.log.ErrorWithID(ctx, "[Controller: Login] Failed to get user: ", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.log.ErrorWithID(ctx, "[Controller: Login] Failed to login user: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.log.InfoWithID(ctx, "[Controller: Login] Successfully logged in user")
	ctx.JSON(http.StatusOK, resp)
}
