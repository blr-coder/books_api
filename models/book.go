package models

import (
	"time"
)

type Book struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"-" gorm:"Column:created_at"`
	UpdatedAt time.Time  `json:"-" gorm:"Column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"Column:deleted_at"`
}
