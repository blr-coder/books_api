package models

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"Column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"Column:deleted_at"`
}
