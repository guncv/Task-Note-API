package containers

import (
	"github.com/guncv/tech-exam-software-engineering/repositories"
)

func (c *Container) RepositoryProvider() {
	if err := c.Container.Provide(repositories.NewTaskRepository); err != nil {
		c.Error = err
	}
}
