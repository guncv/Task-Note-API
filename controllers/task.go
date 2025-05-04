package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
