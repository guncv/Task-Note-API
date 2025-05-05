package main

import (
	"github.com/guncv/tech-exam-software-engineering/containers"
	_ "github.com/guncv/tech-exam-software-engineering/docs"
)

// @title           Tech Exam Hugeman Assignment API
// @version         1.0
// @description     This is a sample Hugeman Assignment backend using Go and Swagger. The project serves as a backend API for managing tasks and users.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	c := containers.NewContainer()
	if err := c.Run().Error; err != nil {
		panic(err)
	}
}
