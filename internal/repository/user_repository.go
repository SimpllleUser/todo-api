package repository

import (
	model "example/todo-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindByID(id uint) (*model.UserModel, error) {
	var user model.UserModel
	err := ur.db.Find(&user, id).Error
	return &user, err
}

func (ur *UserRepository) FindByName(name string) (*model.UserModel, error) {
	var user model.UserModel
	err := ur.db.Where("name LIKE ?", "%"+name+"%").First(&user).Error
	return &user, err
}

func (ur *UserRepository) FindByLogin(login string) (*model.UserModel, error) {
	var user model.UserModel
	err := ur.db.Where("login LIKE ?", "%"+login+"%").First(&user).Error
	return &user, err
}

func (ur *UserRepository) Create(user *model.UserModel) (*model.UserModel, error) {
	err := ur.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
