package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Login     string `json:"login" gorm:"unique"`
	Name      string `json:"name" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) FindById(id uint) (*UserModel, error) {
	var user UserModel
	err := u.db.Find(&user, id).Error
	return &user, err
}

func (u *UserService) FindByName(name string) (*UserModel, error) {
	var user UserModel
	err := u.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (u *UserService) FindByLogin(login string) (*UserModel, error) {
	var user UserModel
	err := u.db.Where("login = ?", login).First(&user).Error
	return &user, err
}

func (u *UserService) Create(user *UserModel) (*UserModel, error) {
	err := u.db.Create(&user)
	if err != nil {
		return nil, err.Error
	}

	return user, nil
}
