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

func TestValidationCreateTodo(t *testing.T) {
	db := setupTestDB(t)
	todoService := NewTodoService(db)

	tests := []struct {
		name           string
		todo           *model.TodoCreateRequest
		validationRule string
		wantErr        bool
	}{
		{
			name: "Valid todo title",
			todo: &model.TodoCreateRequest{
				Title: "Todo valid title",
			},
			wantErr: false,
		},
		{
			name: "Empty todo title",
			todo: &model.TodoCreateRequest{
				Title: "",
			},
			validationRule: "required",
			wantErr:        true,
		},
		{
			name: "Min lens todo title",
			todo: &model.TodoCreateRequest{
				Title: "Ab",
			},
			validationRule: "minLen",
			wantErr:        true,
		},
		{
			name: "Trim todo title",
			todo: &model.TodoCreateRequest{
				Title: "   ",
			},
			validationRule: "required",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := todoService.Create(tt.todo)
			if tt.wantErr {
				if tt.validationRule == "" {
					t.Fatalf("validationRule must be specified for test case: %s", tt.name)
				}
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.validationRule)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	db := setupTestDB(t)
	todoService := NewTodoService(db)

	todo := &model.TodoCreateRequest{
		Title:       "Test item todo title",
		Description: "Test item todo description",
	}

	err := todoService.Create(todo)
	require.NoError(t, err)

	updateTodo := &model.TodoModel{
		ID:          1,
		Title:       "",
		Description: "Test item todo description updated",
		Completed:   true,
	}

	err = todoService.Update(updateTodo)
	require.NoError(t, err)

	var result model.TodoModel
	require.NoError(t, db.First(&result, 1).Error)

	assert.Equal(t, updateTodo.Title, result.Title)
	assert.Equal(t, updateTodo.Description, result.Description)
	assert.Equal(t, updateTodo.Completed, result.Completed)

}
