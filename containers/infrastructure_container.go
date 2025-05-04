package containers

import (
	"github.com/guncv/tech-exam-software-engineering/config"
	"github.com/guncv/tech-exam-software-engineering/infras/database"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/infras/server"
)

func (c *Container) InfraStructureProvider() {

	c.Container.Provide(func(cfg *config.AppConfig) *log.Logger {
		return log.Initialize(cfg.AppEnv)
	})

	c.Container.Provide(database.ConnectPostgres)

	if err := c.Container.Provide(func(cfg *config.AppConfig) *server.GinServer {
		return server.NewGinServer(cfg, c.Container)
	}); err != nil {
		c.Error = err
	}
}
