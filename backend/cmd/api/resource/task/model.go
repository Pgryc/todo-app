package task

import (
	"time"

	"github.com/google/uuid"
)

type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	CreatedDate   string `json:"created_date"`
	CompletedDate string `json:"completed_date"`
	DeletedDate   string `json:"deleted_date"`
}

type Form struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Task struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Description   string
	CreatedDate   time.Time
	CompletedDate time.Time
	DeletedDate   time.Time
}

type Tasks []*Task
