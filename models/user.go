package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;column:id;primaryKey" json:"id"`
	Email     string    `gorm:"column:email;type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"column:password_hash;type:text;not null" json:"-"`
	FirstName string    `gorm:"column:first_name;type:varchar(100)" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:varchar(100)" json:"last_name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;default:now()" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
