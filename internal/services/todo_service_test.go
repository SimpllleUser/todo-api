package service

import (
	"context"
	model "example/todo-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.TodoModel{}, &model.UserModel{})
	require.NoError(t, err)

	return db
}

func TestCreateTodo(t *testing.T) {
	ctx := context.Background()
	db := setupTestDB(t)
	todoService := NewTodoService(db)

	todo := &model.TodoCreateRequest{
		Title:       "Test title",
		Description: "Test description",
	}

	err := todoService.Create(todo)
	require.NoError(t, err)

	var result model.TodoModel
	err = db.WithContext(ctx).First(&result, "title = ?", "Test title").Error

	require.NoError(t, err)

	assert.Equal(t, "Test title", result.Title)
	assert.Equal(t, "Test description", result.Description)
	assert.False(t, result.Completed)

}
