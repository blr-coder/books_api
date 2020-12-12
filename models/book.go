package models

import (
	"time"
)

type Book struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"Column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"Column:deleted_at"`
}
