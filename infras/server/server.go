// server/server.go
package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/guncv/tech-exam-software-engineering/config"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/infras/routes"
	"github.com/guncv/tech-exam-software-engineering/utils"
	"go.uber.org/dig"
)

type GinServer struct {
	Router     *gin.Engine
	AppConfig  *config.AppConfig
	TokenMaker utils.IPasetoMaker
}

func (s *GinServer) Start() error {
	addr := fmt.Sprintf(":%s", s.AppConfig.AppPort)
	return s.Router.Run(addr)
}

func NewGinServer(c *config.Config, diContainer *dig.Container) *GinServer {
	router := gin.Default()

	RegisterCustomValidations()

	log := log.Initialize(c.AppConfig.AppEnv)
	tokenMaker, err := utils.NewPasetoMaker(c, utils.NewPayloadConstruct(c, log))
	if err != nil {
		panic(err)
	}

	s := &GinServer{
		Router:     router,
		AppConfig:  &c.AppConfig,
		TokenMaker: tokenMaker,
	}

	routes.RegisterRoutes(router, diContainer, s.TokenMaker, log)
	return s
}

func RegisterCustomValidations() {
	// Register custom validation for TaskStatus
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("taskstatus", func(fl validator.FieldLevel) bool {
			status := fl.Field().Interface().(constants.TaskStatus)
			return status == constants.TaskStatusPending || status == constants.TaskStatusCompleted
		})
	}

	// Register custom validation for " " empty string
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
			// Trim whitespace and check if anything is left
			return strings.TrimSpace(fl.Field().String()) != ""
		})
	}
}
