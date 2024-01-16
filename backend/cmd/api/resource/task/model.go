package task

import (
	"time"

	"github.com/google/uuid"
)

type DTO struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Created_date   string `json:"created_date"`
	Completed_date string `json:"completed_date"`
	Deleted_date   string `json:"deleted_date"`
}

type Form struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Item struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Description   string
	CreatedDate   time.Time
	CompletedDate time.Time
	DeletedDate   time.Time
}

type Items []*Item
