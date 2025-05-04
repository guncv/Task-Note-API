package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guncv/tech-exam-software-engineering/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

func RegisterRoutes(e *gin.Engine, c *dig.Container) {

	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	e.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := c.Invoke(func(h *controllers.TaskController) {
		api := e.Group("/api/v1")
		api.GET("/health", h.HealthCheck)

		taskRoutes(api, h)
	}); err != nil {
		panic(err)
	}
}

func taskRoutes(eg *gin.RouterGroup, taskController *controllers.TaskController) {
	tasks := eg.Group("/tasks")
	tasks.POST("", taskController.CreateTask)
}

func userRoutes(eg *gin.RouterGroup, userController *controllers.UserController) {
	// users := eg.Group("/users")
}
