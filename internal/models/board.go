package model

import (
	"time"
)

type BoardBaseModel struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Title     string      `json:"title" gorm:"not null"`
	OwnerID   uint        `json:"owner_id" gorm:"not null"`
	Users     []UserModel `json:"users,omitempty" gorm:"many2many:board_users"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BoardModel struct {
	BoardBaseModel
	OwnerID uint `json:"owner_id" gorm:"not null"`
}

type BoardCreateRequest struct {
	BoardBaseModel
}

func (BoardModel) TableName() string {
	return "boards"
}

// package model

// type TodoBaseFields struct {
// 	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
// 	Title       string `gorm:"not null" json:"title" validate:"required|minLen:3" filter:"trim" message:"required:{field} is required" label:"Title" example:"Buy groceries"`
// 	Description string `gorm:"null" json:"description" example:"Milk, eggs, bread"`
// 	Completed   bool   `gorm:"default:false" json:"completed" example:"false"`
// 	TodoProjectId
// }

// type TodoProjectId struct {
// 	ProjectId uint `gorm:"not null" json:"project_id" example:"1"`
// }

// type TodoModel struct {
// 	TodoBaseFields
// 	UserId uint `gorm:"not null" json:"user_id" example:"1"`
// }
// type TodoCreateRequest struct {
// 	TodoBaseFields
// }

// func (TodoModel) TableName() string {
// 	return "todos"
// }
