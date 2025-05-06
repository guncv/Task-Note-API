package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/infras/log"
	"github.com/guncv/tech-exam-software-engineering/services"
	"github.com/guncv/tech-exam-software-engineering/utils"
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

// @Tags Health Check
// @Summary Health Check
// @Description Returns status of the service
// @Success 200 {object} entities.GetHealthUserResponse "Health check successful"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/health [get]
// HealthCheck handles the health check endpoint
func (h *TaskController) HealthCheck(c *gin.Context) {
	reqCtx := c.Request.Context()
	h.log.DebugWithID(reqCtx, "[Controller: HealthCheck] Called")
	response, err := h.service.HealthCheck(reqCtx)
	if err != nil {
		h.log.ErrorWithID(reqCtx, "[Controller: HealthCheck]: Failed to perform health check", err)
		utils.ErrorResponse(c, err)
		return
	}

	h.log.InfoWithID(reqCtx, "[Controller: HealthCheck]: Health check successful")
	c.JSON(http.StatusOK, response)
}

// @Tags Tasks
// @Summary Create Task
// @Description Create a new task
// @Accept multipart/form-data
// @Param title formData string true "Title"
// @Param description formData string false "Description"
// @Param status formData string true "Status"
// @Param date formData string true "Date (RFC3339 format)"
// @Param image formData file false "Optional base64 image upload"
// @Security BearerAuth
// @Success 200 {object} entities.CreateTaskResponse
// @Failure 400 {object} entities.ErrExampleInvalidRequest "Invalid request body"
// @Failure 401 {object} entities.ErrExampleUnauthorized "Unauthorized"
// @Failure 409 {object} entities.ErrExampleTaskAlreadyExists "Task already exists"
// @Failure 404 {object} entities.ErrExampleTaskNotFound "Task not found"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/tasks [post]
func (h *TaskController) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: CreateTask] Called")

	var req entities.CreateTaskRequest
	// Bind request
	if err := c.ShouldBind(&req); err != nil {
		detail := utils.ValidateCreateTaskInput(req)
		h.log.ErrorWithID(ctx, "[Controller: CreateTask]: Failed to bind request", err)
		utils.ErrorResponse(c, constants.ErrInvalidRequestBody, detail)
		return
	}

	// Create task
	response, err := h.service.CreateTask(ctx, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: CreateTask]: Failed to create task", err)
		utils.ErrorResponse(c, err)
		return
	}

	h.log.InfoWithID(ctx, "[Controller: CreateTask]: Task created successfully")
	c.JSON(http.StatusOK, response)
}

// @Tags Tasks
// @Summary Get Task
// @Description Get a task by ID
// @Accept json
// @Param id path string true "Task ID"
// @Security BearerAuth
// @Success 200 {object} entities.GetTaskResponse
// @Failure 401 {object} entities.ErrExampleUnauthorized "Unauthorized"
// @Failure 404 {object} entities.ErrExampleTaskNotFound "Task not found"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/tasks/{id} [get]
func (h *TaskController) GetTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: GetTask] Called")

	// Get task id from path
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, constants.ErrInvalidRequestParam)
		return
	}

	// Get task
	response, err := h.service.GetTask(ctx, id)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: GetTask]: Failed to get task", err)
		utils.ErrorResponse(c, err)
		return
	}

	h.log.InfoWithID(ctx, "[Controller: GetTask]: Task retrieved successfully")
	c.JSON(http.StatusOK, response)
}

// @Tags Tasks
// @Summary Update Task
// @Description Update a task by ID
// @Accept multipart/form-data
// @Param id path string true "Task ID"
// @Param title formData string false "Title"
// @Param description formData string false "Description"
// @Param status formData string false "Status"
// @Param image formData file false "Image"
// @Security BearerAuth
// @Success 200 {object} entities.UpdateTaskResponse "Task updated successfully"
// @Failure 400 {object} entities.ErrExampleInvalidRequest "Invalid request body"
// @Failure 401 {object} entities.ErrExampleUnauthorized "Unauthorized"
// @Failure 404 {object} entities.ErrExampleTaskNotFound "Task not found"
// @Failure 409 {object} entities.ErrExampleTaskAlreadyExists "Task already exists"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/tasks/{id} [put]
func (h *TaskController) UpdateTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: UpdateTask] Called")

	// Get task id from path
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, constants.ErrInvalidRequestParam)
		return
	}

	// Bind request
	var req entities.UpdateTaskRequest
	if err := c.ShouldBind(&req); err != nil {
		detail := utils.ValidateUpdateTaskInput(req)
		h.log.ErrorWithID(ctx, "[Controller: UpdateTask]: Failed to bind request", err)
		utils.ErrorResponse(c, constants.ErrInvalidRequestBody, detail)
		return
	}

	// Update task
	response, err := h.service.UpdateTask(ctx, id, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: UpdateTask]: Failed to update task", err)
		utils.ErrorResponse(c, err)
		return
	}

	h.log.InfoWithID(ctx, "[Controller: UpdateTask]: Task updated successfully")
	c.JSON(http.StatusOK, response)
}

// @Tags Tasks
// @Summary Delete Task
// @Description Delete a task by ID
// @Accept json
// @Param id path string true "Task ID"
// @Security BearerAuth
// @Success 200 {object} nil "Task deleted successfully"
// @Failure 401 {object} entities.ErrExampleUnauthorized "Unauthorized"
// @Failure 404 {object} entities.ErrExampleTaskNotFound "Task not found"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/tasks/{id} [delete]
func (h *TaskController) DeleteTask(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: DeleteTask] Called")

	// Get task id from path
	id := c.Param("id")
	if id == "" {
		utils.ErrorResponse(c, constants.ErrInvalidRequestParam)
		return
	}

	// Delete task
	if err := h.service.DeleteTask(ctx, id); err != nil {
		h.log.ErrorWithID(ctx, "[Controller: DeleteTask]: Failed to delete task", err)
		utils.ErrorResponse(c, err)
		return
	}

	h.log.InfoWithID(ctx, "[Controller: DeleteTask]: Task deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// @Tags Tasks
// @Summary Get All Tasks
// @Description Get all tasks with optional search, sort, and pagination
// @Accept json
// @Produce json
// @Param search query string false "Search by title or description"
// @Param sort_by query string false "Sort by field: title, created_at, status"
// @Param order query string false "Order: asc or desc"
// @Param limit query int false "Number of items per page (default 20)"
// @Param offset query int false "Offset (default 0)"
// @Security BearerAuth
// @Success 200 {object} entities.GetAllTasksResponse "Tasks retrieved successfully"
// @Failure 400 {object} entities.ErrExampleInvalidRequest "Invalid query params"
// @Failure 401 {object} entities.ErrExampleUnauthorized "Unauthorized"
// @Failure 500 {object} entities.ErrExampleInternalError "Internal server error"
// @Router /api/v1/tasks [get]
func (h *TaskController) GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.DebugWithID(ctx, "[Controller: GetAllTasks] Called")

	var req entities.GetAllTasksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		detail := utils.ValidateGetAllTasksInput(req)
		h.log.ErrorWithID(ctx, "[Controller: GetAllTasks]: Invalid query params", err)
		utils.ErrorResponse(c, constants.ErrInvalidQueryRequestParam, detail)
		return
	}

	response, err := h.service.GetAllTasks(ctx, &req)
	if err != nil {
		h.log.ErrorWithID(ctx, "[Controller: GetAllTasks]: Failed to get all tasks", err)
		utils.ErrorResponse(c, err)
		return
	}

	h.log.InfoWithID(ctx, "[Controller: GetAllTasks]: Tasks retrieved successfully")
	c.JSON(http.StatusOK, response)
}
