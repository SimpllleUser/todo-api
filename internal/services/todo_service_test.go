package service

import (
	"context"
	model "example/todo-api/internal/models"
	test_utils "example/todo-api/internal/testutils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func seedTodos(t *testing.T, service *TodoService, count int) []model.TodoModel {
	todos := make([]model.TodoModel, 0, count)

	for index := 0; index < count; index++ {
		todoReq := model.TodoCreateRequest{
			Title:       fmt.Sprintf("Todo title #%d", index),
			Description: fmt.Sprintf("Todo description #%d", index),
		}

		createdTodo, err := service.Create(&todoReq)
		require.NoError(t, err)

		todos = append(todos, model.TodoModel{
			ID:          createdTodo.ID,
			Title:       createdTodo.Title,
			Description: createdTodo.Description,
			Completed:   createdTodo.Completed,
		})

	}

	return todos
}

func TestCreateTodo(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	todo := &model.TodoCreateRequest{
		Title:       "Test title",
		Description: "Test description",
	}

	createdTodo, err := todoService.Create(todo)
	require.NoError(t, err)

	var result model.TodoModel
	err = db.WithContext(ctx).First(&result, "title = ?", "Test title").Error

	require.NoError(t, err)

	assert.Equal(t, createdTodo.Title, result.Title)
	assert.Equal(t, createdTodo.Description, result.Description)
	assert.False(t, result.Completed)

}

func TestValidationCreateTodo(t *testing.T) {
	t.Parallel()
	db := test_utils.SetupTestDB(t)
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
			_, err := todoService.Create(tt.todo)
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
	t.Parallel()

	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	todo := &model.TodoCreateRequest{
		Title:       "Test item todo title",
		Description: "Test item todo description",
	}

	createdTodo, err := todoService.Create(todo)
	require.NoError(t, err)

	updateTodo := &model.TodoModel{
		ID:          createdTodo.ID,
		Title:       "Test item todo title updated",
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

func TestValidationUpdateTodo(t *testing.T) {

	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	tests := []struct {
		name           string
		todo           *model.TodoModel
		validationRule string
		wantErr        bool
	}{
		{
			name: "Valid todo title",
			todo: &model.TodoModel{
				ID:    1,
				Title: "Todo valid title",
			},
			wantErr: false,
		},
		{
			name: "Empty todo title",
			todo: &model.TodoModel{
				ID:    1,
				Title: "",
			},
			validationRule: "required",
			wantErr:        true,
		},
		{
			name: "Min lens todo title",
			todo: &model.TodoModel{
				ID:    1,
				Title: "Ab",
			},
			validationRule: "minLen",
			wantErr:        true,
		},
		{
			name: "Trim todo title",
			todo: &model.TodoModel{
				ID:    1,
				Title: "   ",
			},
			validationRule: "required",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := todoService.Update(tt.todo)
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

func TestFindByIdTodo(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	todo := &model.TodoCreateRequest{
		Title: "Test title",
	}

	_, err := todoService.Create(todo)
	require.NoError(t, err)

	const todoID = uint(1)

	found, err := todoService.GetById(todoID)

	require.NoError(t, err)

	assert.Equal(t, todoID, found.ID)
	assert.Equal(t, todo.Title, found.Title)
	assert.Equal(t, todo.Description, found.Description)
	assert.Equal(t, todo.Completed, found.Completed)

}

func TestFindByTitleOneMoreTodos(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	todos := seedTodos(t, todoService, 10)

	todoFromList := todos[:2]

	foundTodos, err := todoService.GetByTitle("Todo")

	assert.Equal(t, foundTodos, todoFromList)
	assert.Len(t, foundTodos, 2)

	require.NoError(t, err)

}

func TestFindByTitleTodos(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	seedTodos := seedTodos(t, todoService, 10)

	currentTodo := seedTodos[5]

	foundTodos, err := todoService.GetByTitle("Todo title #5")

	assert.Equal(t, foundTodos[0], currentTodo)

	require.NoError(t, err)

}

func TestFindByTitleTodosWithEmptyResult(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	seedTodos(t, todoService, 10)

	foundTodos, err := todoService.GetByTitle("Some test random text qwerty")

	assert.Len(t, foundTodos, 0)

	require.NoError(t, err)

}

// 21asdasd

func TestGetAll(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	seededTodos := seedTodos(t, todoService, 10)

	foundTodos, err := todoService.GetAll()

	assert.Len(t, foundTodos, 10)

	for index, seededTodo := range seededTodos {
		assert.Equal(t, seededTodo, *foundTodos[index])
	}

	require.NoError(t, err)

}

func TestDelete(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	seedTodos(t, todoService, 3)

	todoListBeforeRemove, err := todoService.GetAll()
	require.NoError(t, err)

	errOnDelete := todoService.Delete(1)
	require.NoError(t, errOnDelete)

	todoListAfterRemove, err := todoService.GetAll()
	require.NoError(t, err)

	assert.Len(t, todoListAfterRemove, len(todoListBeforeRemove)-1)

}

func TestDeleteWithUnrealId(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTodoService(db)

	seedTodos(t, todoService, 3)

	todoListBeforeRemove, err := todoService.GetAll()
	require.NoError(t, err)

	errOnDelete := todoService.Delete(123)
	require.NoError(t, errOnDelete)

	assert.Equal(t, errOnDelete, nil)

	todoListAfterRemove, err := todoService.GetAll()
	require.NoError(t, err)

	assert.Len(t, todoListBeforeRemove, len(todoListAfterRemove))

}
