package service

import (
	model "example/todo-api/internal/models"
	"example/todo-api/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) FindById(id uint) (*model.UserModel, error) {
	return u.UserRepository.FindByID(id)
}

func (u *UserService) FindByName(name string) (*model.UserModel, error) {
	return u.UserRepository.FindByName(name)
}

func (u *UserService) FindByLogin(login string) (*model.UserModel, error) {
	return u.UserRepository.FindByLogin(login)
}

func (u *UserService) Create(user *model.UserModel) (*model.UserModel, error) {
	return u.UserRepository.Create(user)
}
