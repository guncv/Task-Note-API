package main

import (
	"github.com/guncv/tech-exam-software-engineering/containers"
	_ "github.com/guncv/tech-exam-software-engineering/docs"
)

// @title       Tech Exam Hugeman Assignment API
// @version     1.0
// @description This is a sample Hugeman Assignment backend using Go and Swagger. The project serves as a backend API for managing tasks and users.
// @host        localhost:8080
// @BasePath    /api/v1
func main() {
	c := containers.NewContainer()
	if err := c.Run().Error; err != nil {
		panic(err)
	}
}
