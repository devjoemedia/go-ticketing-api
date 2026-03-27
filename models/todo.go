package models

import (
	"time"

	"gorm.io/gorm"
)

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	IsCompleted bool   `json:"is_completed" validate:"omitempty"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	IsCompleted bool   `json:"is_completed" validate:"omitempty"`
}

type Todo struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
