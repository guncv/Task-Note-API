package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/services"
)

type TaskController struct {
	service services.ITaskService
	log     *log.Logger
}

func NewTaskController(service services.ITaskService, log *log.Logger) *TaskController {
	return &TaskController{
		service: service,
		log:     log,
	}
}

// @Summary Health Check
// @Description Returns status
// @Success 200 {object} entities.GetHealthUserResponse
// @Tags Health
// @Router /v1/health [get]
// HealthCheck handles the health check endpoint
func (h *TaskController) HealthCheck(c *gin.Context) {
	reqCtx := c.Request.Context()
	h.log.DebugWithID(reqCtx, "[Controller: HealthCheck] Called")
	response, err := h.service.HealthCheck(reqCtx)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Controller: HealthCheck]: Failed to perform health check", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to health check"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Create Task
// @Description Create a new task
// @Accept multipart/form-data
// @Param title formData string true "Title"
// @Param description formData string false "Description"
// @Param status formData string true "Status"
// @Param date formData string true "Date"
// @Param image formData file false "Image"
// @Success 200 {object} entities.CreateTaskResponse
// @Router /v1/tasks [post]
func (h *TaskController) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: CreateTask] Called")

	var req entities.CreateTaskRequest
	// Bind request
	if err := c.ShouldBind(&req); err != nil {
		h.log.ErrorWithID(ctx, "[Controller: CreateTask]: Failed to bind request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	response, err := h.service.CreateTask(ctx, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: CreateTask]: Failed to create task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	h.log.InfoWithID(ctx, "[Controller: CreateTask]: Task created successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": response})
}
