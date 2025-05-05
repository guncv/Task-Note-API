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
	Description *string               `form:"description" binding:"omitempty"`
	Status      constants.TaskStatus  `form:"status" binding:"required,taskstatus"`
	Image       *multipart.FileHeader `form:"image" binding:"omitempty"`
}

type CreateTaskResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	CreatedAt   string  `json:"created_at"`
}

type GetTaskResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	CreatedAt   string  `json:"created_at"`
}

type UpdateTaskRequest struct {
	Title       string                `form:"title" binding:"omitempty,max=50,notblank"`
	Description *string               `form:"description" binding:"omitempty"`
	Status      constants.TaskStatus  `form:"status" binding:"omitempty,taskstatus"`
	Image       *multipart.FileHeader `form:"image" binding:"omitempty"`
}

type UpdateTaskResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Status      string  `json:"status"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	CreatedAt   string  `json:"created_at"`
}

type GetAllTasksRequest struct {
	Search string `form:"search"`
	SortBy string `form:"sort_by" binding:"omitempty,oneof=title created_at status"`
	Order  string `form:"order" binding:"omitempty,oneof=asc desc"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int    `form:"offset" binding:"omitempty,min=1"`
}

type GetAllTasksResponse struct {
	Total int               `json:"total"`
	Tasks []GetTaskResponse `json:"tasks"`
}
