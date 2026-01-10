package model

import (
	"time"
)

type UserModel struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Login     string `json:"login" gorm:"unique"`
	Name      string `json:"name" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}

type UserCreateRequest struct {
	User  UserModel `json:"user"`
	Token string    `json:"token" binding:"required"`
}
