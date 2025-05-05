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
	"github.com/guncv/tech-exam-software-engineering/infras/routes"
	"go.uber.org/dig"
)

type GinServer struct {
	Router    *gin.Engine
	AppConfig *config.AppConfig
}

func (s *GinServer) Start() error {
	addr := fmt.Sprintf(":%s", s.AppConfig.AppPort)
	return s.Router.Run(addr)
}

func NewGinServer(c *config.Config, diContainer *dig.Container) *GinServer {
	router := gin.Default()

	RegisterCustomValidations()

	s := &GinServer{
		Router:    router,
		AppConfig: &c.AppConfig,
	}

	routes.RegisterRoutes(router, diContainer)
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
