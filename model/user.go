package model

import "gorm.io/gorm"

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

func (u *UserService) FindById(id uint) (*User, error) {
	var user User
	err := u.db.Find(&user, id).Error
	return &user, err
}

func (u *UserService) FindByName(name string) (*User, error) {
	var user User
	err := u.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (u *UserService) FindByLogin(login string) (*User, error) {
	var user User
	err := u.db.Where("login = ?", login).First(&user).Error
	return &user, err
}

func (u *UserService) Create(user *User) (*User, error) {
	err := u.db.Create(&user)
	if err != nil {
		return nil, err.Error
	}

	return user, nil
}
