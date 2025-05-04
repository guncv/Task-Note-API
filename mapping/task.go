package mapping

import (
	"time"

	"github.com/google/uuid"
	"github.com/guncv/tech-exam-software-engineering/entities"
	"github.com/guncv/tech-exam-software-engineering/models"
)

func MapCreateTaskRequestToTask(req *entities.CreateTaskRequest, base64Image string, id uuid.UUID) *models.Task {
	return &models.Task{
		ID:          id,
		Title:       req.Title,
		Status:      string(req.Status),
		Image:       &base64Image,
		Description: &req.Description,
	}
}

func MapCreateTaskResponseToTask(task *models.Task, message string) *entities.CreateTaskResponse {
	return &entities.CreateTaskResponse{
		Message:     message,
		ID:          task.ID.String(),
		Title:       task.Title,
		Status:      task.Status,
		Image:       *task.Image,
		Description: *task.Description,
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
	}
}
