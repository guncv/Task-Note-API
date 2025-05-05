package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guncv/tech-exam-software-engineering/controllers"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/middleware"
	"github.com/guncv/tech-exam-software-engineering/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

func RegisterRoutes(e *gin.Engine, c *dig.Container, tokenMaker utils.IPasetoMaker, log *log.Logger) {

	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	e.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := c.Invoke(func(
		taskController *controllers.TaskController,
		userController *controllers.UserController,
	) {
		api := e.Group("/api/v1")
		api.GET("/health", taskController.HealthCheck)

		userRoutes(api, userController)

		// Auth Middleware Routes
		authRoutes := api.Group("/").Use(middleware.AuthMiddleware(tokenMaker, log))

		taskRoutes(authRoutes.(*gin.RouterGroup), taskController)
	}); err != nil {
		panic(err)
	}
}

func taskRoutes(eg *gin.RouterGroup, taskController *controllers.TaskController) {
	tasks := eg.Group("/tasks")
	tasks.POST("", taskController.CreateTask)
	tasks.GET("", taskController.GetAllTasks)
	tasks.GET("/:id", taskController.GetTask)
	tasks.PUT("/:id", taskController.UpdateTask)
	tasks.DELETE("/:id", taskController.DeleteTask)
}

func userRoutes(eg *gin.RouterGroup, userController *controllers.UserController) {
	users := eg.Group("/users")
	users.POST("", userController.Register)
	users.POST("/login", userController.Login)
}
