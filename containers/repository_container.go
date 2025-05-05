package containers

import (
	"github.com/guncv/tech-exam-software-engineering/repositories"
	"github.com/guncv/tech-exam-software-engineering/utils"
)

func (c *Container) RepositoryProvider() {
	if err := c.Container.Provide(repositories.NewTaskRepository); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(repositories.NewUserRepository); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(utils.NewPayload); err != nil {
		c.Error = err
	}
}
