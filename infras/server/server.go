// server/server.go
package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guncv/tech-exam-software-engineering/config"
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

func NewGinServer(c *config.AppConfig, diContainer *dig.Container) *GinServer {
	router := gin.Default()

	s := &GinServer{
		Router:    router,
		AppConfig: c,
	}

	routes.RegisterRoutes(router, diContainer)

	return s
}
