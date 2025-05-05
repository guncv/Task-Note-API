package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;column:id;primaryKey" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;column:user_id;not null" validate:"required" json:"user_id"`
	Title       string    `gorm:"column:title;type:varchar(100);not null" validate:"required" json:"title"`                          // Title of the task
	Description *string   `gorm:"column:description;type:text" json:"description,omitempty"`                                         // Optional description
	Image       *string   `gorm:"column:image;type:text" json:"image,omitempty"`                                                     // Base64-encoded image string
	Status      string    `gorm:"column:status;type:varchar(20);not null;check:status IN ('IN_PROGRESS','COMPLETED')" json:"status"` // Status: IN_PROGRESS or COMPLETED
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamptz;not null;default:now()" json:"created_at"`
}

// TableName overrides the default table name used by GORM
func (Task) TableName() string {
	return "tasks"
}
