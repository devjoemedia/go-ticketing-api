package models

import (
	"time"

	"gorm.io/gorm"
)

type CreateTicketRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Assignee    string `json:"assignee"`
	Attachments string `json:"attachments"`
}

type UpdateTicketRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Assignee    string `json:"assignee"`
	Comments    string `json:"comments"`
	Attachments string `json:"attachments"`
	HistoryLogs string `json:"history_logs"`
}

type Ticket struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Reporter    string `json:"reporter"`
	Assignee    string `json:"assignee"`
	Comments    string `json:"comments"`
	Attachments string `json:"attachments"`
	HistoryLogs string `json:"history_logs"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
