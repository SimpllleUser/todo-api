package service

import (
	model "example/todo-api/internal/models"
	"example/todo-api/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) FindById(id uint) (*model.UserModel, error) {
	return u.userRepository.FindByID(id)
}

func (u *UserService) FindByName(name string) (*model.UserModel, error) {
	return u.userRepository.FindByName(name)
}

func (u *UserService) FindByLogin(login string) (*model.UserModel, error) {
	return u.userRepository.FindByLogin(login)
}

func (u *UserService) Create(user *model.UserModel) (*model.UserModel, error) {
	return u.userRepository.Create(user)
}
