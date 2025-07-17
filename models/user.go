package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
// @Description User account information
type User struct {
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name" gorm:"not null" example:"John Doe"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null" example:"john@example.com"`
	Age       int            `json:"age" example:"30"`
}

// CreateUserRequest represents the request body for creating a user
// @Description Request body for creating a new user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
	Age   int    `json:"age" binding:"min=1,max=120" example:"30"`
}

// UpdateUserRequest represents the request body for updating a user
// @Description Request body for updating an existing user
type UpdateUserRequest struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" binding:"omitempty,email" example:"john@example.com"`
	Age   int    `json:"age" binding:"min=1,max=120" example:"30"`
}
