package containers

import "github.com/guncv/tech-exam-software-engineering/controllers"

func (c *Container) ControllerProvider() {
	if err := c.Container.Provide(controllers.NewTaskController); err != nil {
		c.Error = err
	}
}
