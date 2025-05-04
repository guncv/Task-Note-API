package containers

import (
	"github.com/guncv/tech-exam-software-engineering/config"
	"github.com/guncv/tech-exam-software-engineering/infras/server"
	"go.uber.org/dig"
)

type Container struct {
	Container *dig.Container
	Error     error
}

func (c *Container) Configure() {
	c.Container = dig.New()

	c.Container.Provide(config.LoadConfig)

	c.ControllerProvider()
	c.ServiceProvider()
	c.RepositoryProvider()
	c.InfraStructureProvider()
}

func (c *Container) Run() *Container {
	if err := c.Container.Invoke(func(s *server.GinServer) {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
	return c
}

func NewContainer() *Container {
	c := &Container{}
	c.Configure()
	return c
}
