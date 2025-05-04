package containers

import (
	"github.com/guncv/tech-exam-software-engineering/services"
)

func (c *Container) ServiceProvider() {
	if err := c.Container.Provide(services.NewTaskService); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(services.NewUserService); err != nil {
		c.Error = err
	}
}
