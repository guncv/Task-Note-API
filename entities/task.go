package entities

import (
	"mime/multipart"

	constants "github.com/guncv/tech-exam-software-engineering/constant"
)

type GetHealthUserRequest struct {
	Service string `json:"service"`
}

type GetHealthUserResponse struct {
	Status string `json:"status"`
}

type CreateTaskRequest struct {
	Title       string                `form:"title" binding:"required,max=50,notblank"`
	Description string                `form:"description"`
	Status      constants.TaskStatus  `form:"status" binding:"required,taskstatus"`
	Image       *multipart.FileHeader `form:"image"`
}

type CreateTaskResponse struct {
	Message     string `json:"message"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Date        string `json:"date"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
