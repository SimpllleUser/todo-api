package service

import (
	model "example/todo-api/internal/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUserTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.UserModel{})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Exec("DELETE FROM users").Error
		require.NoError(t, err)
	})
	return db
}

func seedUsers(t *testing.T, service *UserService, count int) []model.UserModel {
	users := make([]model.UserModel, 0, count)

	for index := 0; index < count; index++ {
		userReq := &model.UserModel{
			Login:    fmt.Sprintf("User-%d", index),
			Password: fmt.Sprintf("UserPass-%d", index),
		}

		createdUser, err := service.Create(userReq)
		require.NoError(t, err)

		users = append(users, model.UserModel{
			ID:       createdUser.ID,
			Login:    createdUser.Login,
			Password: createdUser.Password,
		})

	}

	return users
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	db := setupUserTestDB(t)
	userService := NewUserService(db)

	userReq := &model.UserModel{
		Login:    "test",
		Password: "testPass",
	}

	createdUser, err := userService.Create(userReq)
	require.Nil(t, err)

	assert.Equal(t, userReq.Login, createdUser.Login)
	assert.Equal(t, userReq.Password, createdUser.Password)
}

func TestFindById(t *testing.T) {
	t.Parallel()
	db := setupUserTestDB(t)
	userService := NewUserService(db)

	seededUsers := seedUsers(t, userService, 10)

	foundUser, err := userService.FindById(3)
	require.NoError(t, err)

	expectedUser := seededUsers[2]

	assert.Equal(t, foundUser.ID, expectedUser.ID)
	assert.Equal(t, foundUser.Login, expectedUser.Login)
	assert.Equal(t, foundUser.Password, expectedUser.Password)
}
