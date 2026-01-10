package model

import (
	"time"
)

type ProjectModel struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Users       []UserModel
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ProjectModel) TableName() string {
	return "projects"
}

type ProjectCreateRequest struct {
	Project ProjectModel
}
