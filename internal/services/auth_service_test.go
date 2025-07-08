package service

import (
	model "example/todo-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	FindByLogin(login string) (*model.UserModel, error)
}

type fakerUserService struct {
	users map[string]*model.UserModel
}

func (f *fakerUserService) FindByLogin(login string) (*model.UserModel, error) {
	user, exist := f.users[login]
	if !exist {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

func Test_User_Auth_Success(t *testing.T) {
	password := "my_pass"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockUser := &model.UserModel{
		ID:       1,
		Login:    "my_login",
		Password: string(hashedPassword),
	}

	fakeService := &fakerUserService{
		users: map[string]*model.UserModel{
			"my_login": mockUser,
		},
	}

	authService := NewAuthService(fakeService)

	authedUser, err := authService.Authenticate("my_login", password)

	require.NoError(t, err)

	assert.Equal(t, authedUser, mockUser)
}

func Test_User_Auth_InvalidPassword(t *testing.T) {
	password := "my_pass"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockUser := &model.UserModel{
		ID:       1,
		Login:    "my_login",
		Password: string(hashedPassword),
	}

	fakeService := &fakerUserService{
		users: map[string]*model.UserModel{
			"my_login": mockUser,
		},
	}

	authService := NewAuthService(fakeService)

	authedUser, err := authService.Authenticate("my_login", "my_invalid_pass")

	require.Error(t, err)

	assert.Nil(t, authedUser)
}

func Test_User_Auth_UserNotFound(t *testing.T) {
	password := "my_pass"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockUser := &model.UserModel{
		ID:       1,
		Login:    "my_login",
		Password: string(hashedPassword),
	}

	fakeService := &fakerUserService{
		users: map[string]*model.UserModel{
			"my_login": mockUser,
		},
	}

	authService := NewAuthService(fakeService)

	authedUser, err := authService.Authenticate("my_invalid_login", password)

	require.Error(t, err)

	assert.Nil(t, authedUser)
}
