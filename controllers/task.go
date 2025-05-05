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

	h.log.InfoWithID(reqCtx, "[Controller: HealthCheck]: Health check successful")
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

	// Create task
	response, err := h.service.CreateTask(ctx, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: CreateTask]: Failed to create task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	h.log.InfoWithID(ctx, "[Controller: CreateTask]: Task created successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": response})
}

// @Summary Get Task
// @Description Get a task by ID
// @Accept json
// @Param id path string true "Task ID"
// @Success 200 {object} entities.GetTaskResponse
// @Router /v1/tasks/{id} [get]
func (h *TaskController) GetTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: GetTask] Called")

	// Get task id from path
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID in path"})
		return
	}

	// Get task
	response, err := h.service.GetTask(ctx, id)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: GetTask]: Failed to get task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	h.log.InfoWithID(ctx, "[Controller: GetTask]: Task retrieved successfully")
	c.JSON(http.StatusOK, response)
}

// @Summary Update Task
// @Description Update a task by ID
// @Accept multipart/form-data
// @Param id path string true "Task ID"
// @Param title formData string false "Title"
// @Param description formData string false "Description"
// @Param status formData string false "Status"
// @Param image formData file false "Image"
// @Success 200 {object} entities.UpdateTaskResponse
// @Router /v1/tasks/{id} [put]
func (h *TaskController) UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: UpdateTask] Called")

	// Get task id from path
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID in path"})
		return
	}

	// Bind request
	var req entities.UpdateTaskRequest
	if err := c.ShouldBind(&req); err != nil {
		h.log.ErrorWithID(ctx, "[Controller: UpdateTask]: Failed to bind request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Update task
	response, err := h.service.UpdateTask(ctx, id, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: UpdateTask]: Failed to update task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	h.log.InfoWithID(ctx, "[Controller: UpdateTask]: Task updated successfully")
	c.JSON(http.StatusOK, response)
}

// @Summary Delete Task
// @Description Delete a task by ID
// @Accept json
// @Param id path string true "Task ID"
// @Success 200 {object} entities.DeleteTaskResponse
// @Router /v1/tasks/{id} [delete]
func (h *TaskController) DeleteTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: DeleteTask] Called")

	// Get task id from path
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID in path"})
		return
	}

	// Delete task
	if err := h.service.DeleteTask(ctx, id); err != nil {
		h.log.ErrorWithID(ctx, "[Controller: DeleteTask]: Failed to delete task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	h.log.InfoWithID(ctx, "[Controller: DeleteTask]: Task deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// @Summary Get All Tasks
// @Description Get all tasks with optional search, sort, and pagination
// @Accept json
// @Produce json
// @Param search query string false "Search by title or description"
// @Param sort_by query string false "Sort by field: title, created_at, status"
// @Param order query string false "Order: asc or desc"
// @Param limit query int false "Number of items per page (default 20)"
// @Param offset query int false "Offset (default 0)"
// @Success 200 {object} entities.GetAllTasksResponse
// @Failure 400 {object} gin.H{"error": "Invalid query parameters"}
// @Failure 500 {object} gin.H{"error": "Failed to get all tasks"}
// @Router /v1/tasks [get]
func (h *TaskController) GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: GetAllTasks] Called")

	var req entities.GetAllTasksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.ErrorWithID(ctx, "[Controller: GetAllTasks]: Invalid query params", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	response, err := h.service.GetAllTasks(ctx, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: GetAllTasks]: Failed to get all tasks", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all tasks"})
		return
	}

	h.log.InfoWithID(ctx, "[Controller: GetAllTasks]: Tasks retrieved successfully")
	c.JSON(http.StatusOK, response)
}
