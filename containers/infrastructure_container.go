package containers

import (
	"github.com/guncv/tech-exam-software-engineering/config"
	"github.com/guncv/tech-exam-software-engineering/infras/database"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/infras/server"
)

func (c *Container) InfraStructureProvider() {

	c.Container.Provide(func(cfg *config.Config) *log.Logger {
		return log.Initialize(cfg.AppConfig.AppEnv)
	})

	c.Container.Provide(database.ConnectPostgres)

	if err := c.Container.Provide(func(cfg *config.Config) *server.GinServer {
		return server.NewGinServer(cfg, c.Container)
	}); err != nil {
		c.Error = err
	}
}
