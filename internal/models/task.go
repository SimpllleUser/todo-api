package model

import "time"

type TaskModel struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	BoardID     uint      `json:"board_id" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskCreateRequest struct {
	Title       string `json:"title" binding:"required,min=3" validate:"required|minLen:3"`
	Description string `json:"description"`
	BoardID     uint   `json:"board_id" binding:"required" validate:"required"`
}

func (TaskModel) TableName() string {
	return "tasks"
}
