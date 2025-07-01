package service

import (
	model "example/todo-api/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) FindById(id uint) (*model.UserModel, error) {
	var user model.UserModel
	err := u.db.Find(&user, id).Error
	return &user, err
}

func (u *UserService) FindByName(name string) (*model.UserModel, error) {
	var user model.UserModel
	err := u.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (u *UserService) FindByLogin(login string) (*model.UserModel, error) {
	var user model.UserModel
	err := u.db.Where("login = ?", login).First(&user).Error
	return &user, err
}

func (u *UserService) Create(user *model.UserModel) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
