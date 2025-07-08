package service

import (
	model "example/todo-api/internal/models"
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
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

func Test_User_Auth_GenerateToken(t *testing.T) {

	os.Setenv("SECRET_KEY", "test_secret")

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
	tokenStr, err := authService.GenerateToken(mockUser.ID)

	require.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("test_secret"), nil
	})

	require.NoError(t, err)

	assert.True(t, token.Valid)

	claims := token.Claims.(jwt.MapClaims)

	assert.EqualValues(t, mockUser.ID, claims["id"])
}
