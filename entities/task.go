package entities

import (
	"mime/multipart"
	"time"

	constants "github.com/guncv/tech-exam-software-engineering/constant"
)

type GetHealthUserResponse struct {
	Status string `json:"status" example:"Healthy"`
}

type CreateTaskRequest struct {
	Title       string                `form:"title" binding:"required,max=100" example:"Task 1"`
	Description *string               `form:"description" binding:"omitempty" example:"Description of task 1"`
	Status      constants.TaskStatus  `form:"status" binding:"required,taskstatus" example:"IN_PROGRESS"`
	Date        time.Time             `form:"date" time_format:"2006-01-02T15:04:05Z07:00" binding:"required" example:"2021-09-01T00:00:00Z"`
	Image       *multipart.FileHeader `form:"image" binding:"omitempty" example:"https://example.com/image.jpg"`
}

type CreateTaskResponse struct {
	ID          string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserID      string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title       string    `json:"title" example:"Task 1"`
	Status      string    `json:"status" example:"IN_PROGRESS"`
	Description *string   `json:"description" example:"Description of task 1"`
	Date        time.Time `json:"date" example:"2021-09-01T00:00:00Z"`
	Image       *string   `json:"image" example:"fqfqf"`
	CreatedAt   time.Time `json:"created_at" example:"2021-09-01T00:00:00Z"`
}

type GetTaskResponse struct {
	ID          string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserID      string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title       string    `json:"title" example:"Task 1"`
	Status      string    `json:"status" example:"IN_PROGRESS"`
	Description *string   `json:"description" example:"Description of task 1"`
	Date        time.Time `json:"date" example:"2021-09-01T00:00:00Z"`
	Image       *string   `json:"image" example:"fqfqf"`
	CreatedAt   time.Time `json:"created_at" example:"2021-09-01T00:00:00Z"`
}

type UpdateTaskRequest struct {
	Title       string                `form:"title" binding:"omitempty,max=100,notblank" example:"Task 1"`
	Description *string               `form:"description" binding:"omitempty" example:"Description of task 1"`
	Status      constants.TaskStatus  `form:"status" binding:"omitempty,taskstatus" example:"IN_PROGRESS"`
	Date        time.Time             `form:"date" time_format:"2006-01-02T15:04:05Z07:00" binding:"omitempty" example:"2021-09-01T00:00:00Z"`
	Image       *multipart.FileHeader `form:"image" binding:"omitempty" example:"https://example.com/image.jpg"`
}

type UpdateTaskResponse struct {
	ID          string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserID      string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title       string    `json:"title" example:"Task 1"`
	Status      string    `json:"status" example:"IN_PROGRESS"`
	Description *string   `json:"description" example:"Description of task 1"`
	Date        time.Time `json:"date" example:"2021-09-01T00:00:00Z"`
	Image       *string   `json:"image" example:"fqfqf"`
	CreatedAt   time.Time `json:"created_at" example:"2021-09-01T00:00:00Z"`
}

type GetAllTasksRequest struct {
	Search string `form:"search" example:"Task 1"`
	SortBy string `form:"sort_by" binding:"omitempty,oneof=title created_at status" example:"title"`
	Order  string `form:"order" binding:"omitempty,oneof=asc desc" example:"asc"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100" example:"10"`
	Offset int    `form:"offset" binding:"omitempty,min=1" example:"0"`
}

type GetAllTasksResponse struct {
	Total int               `json:"total" example:"1"`
	Tasks []GetTaskResponse `json:"tasks"`
}
