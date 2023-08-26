package models

import (
	"time"
)

type Todo struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title" gorm:"text;not null"`
	Completed *bool      `json:"completed" gorm:"default:false"`
}
