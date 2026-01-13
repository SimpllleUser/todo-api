package model

import "time"

type BoardModel struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Title     string      `json:"title" gorm:"not null" validate:"required"`
	OwnerID   uint        `json:"owner_id" gorm:"not null"`
	Users     []UserModel `json:"users,omitempty" gorm:"many2many:board_users;"`
	Tasks     []TaskModel `json:"tasks,omitempty" gorm:"foreignKey:BoardID"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type BoardCreateRequest struct {
	Title string `json:"title" binding:"required" validate:"required"`
}

func (BoardModel) TableName() string {
	return "boards"
}
