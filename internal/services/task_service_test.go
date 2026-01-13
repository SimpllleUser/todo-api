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

func seedTasks(t *testing.T, service *TaskService, count int) []model.TaskModel {
	todos := make([]model.TaskModel, 0, count)

	for index := 0; index < count; index++ {
		todoReq := model.TaskCreateRequest{
			Title:       fmt.Sprintf("Task title #%d", index),
			Description: fmt.Sprintf("Task description #%d", index),
		}

		createdTask, err := service.Create(&todoReq)
		require.NoError(t, err)

		todos = append(todos, model.TaskModel{
			ID:          createdTask.ID,
			Title:       createdTask.Title,
			Description: createdTask.Description,
			Completed:   createdTask.Completed,
		})

	}

	return todos
}

func TestCreateTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	todo := &model.TaskCreateRequest{
		Title:       "Test title",
		Description: "Test description",
	}

	createdTask, err := todoService.Create(todo)
	require.NoError(t, err)

	var result model.TaskModel
	err = db.WithContext(ctx).First(&result, "title = ?", "Test title").Error

	require.NoError(t, err)

	assert.Equal(t, createdTask.Title, result.Title)
	assert.Equal(t, createdTask.Description, result.Description)
	assert.False(t, result.Completed)

}

func TestValidationCreateTask(t *testing.T) {
	t.Parallel()
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	tests := []struct {
		name           string
		todo           *model.TaskCreateRequest
		validationRule string
		wantErr        bool
	}{
		{
			name: "Valid todo title",
			todo: &model.TaskCreateRequest{
				Title: "Task valid title",
			},
			wantErr: false,
		},
		{
			name: "Empty todo title",
			todo: &model.TaskCreateRequest{
				Title: "",
			},
			validationRule: "required",
			wantErr:        true,
		},
		{
			name: "Min lens todo title",
			todo: &model.TaskCreateRequest{
				Title: "Ab",
			},
			validationRule: "minLen",
			wantErr:        true,
		},
		{
			name: "Trim todo title",
			todo: &model.TaskCreateRequest{
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

func TestUpdateTask(t *testing.T) {
	t.Parallel()

	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	todo := &model.TaskCreateRequest{
		Title:       "Test item todo title",
		Description: "Test item todo description",
	}

	createdTask, err := todoService.Create(todo)
	require.NoError(t, err)

	updateTask := &model.TaskModel{
		ID:          createdTask.ID,
		Title:       "Test item todo title updated",
		Description: "Test item todo description updated",
		Completed:   true,
	}

	err = todoService.Update(updateTask)
	require.NoError(t, err)

	var result model.TaskModel
	require.NoError(t, db.First(&result, 1).Error)

	assert.Equal(t, updateTask.Title, result.Title)
	assert.Equal(t, updateTask.Description, result.Description)
	assert.Equal(t, updateTask.Completed, result.Completed)

}

func TestValidationUpdateTask(t *testing.T) {

	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	tests := []struct {
		name           string
		todo           *model.TaskModel
		validationRule string
		wantErr        bool
	}{
		{
			name: "Valid todo title",
			todo: &model.TaskModel{
				ID:    1,
				Title: "Task valid title",
			},
			wantErr: false,
		},
		{
			name: "Empty todo title",
			todo: &model.TaskModel{
				ID:    1,
				Title: "",
			},
			validationRule: "required",
			wantErr:        true,
		},
		{
			name: "Min lens todo title",
			todo: &model.TaskModel{
				ID:    1,
				Title: "Ab",
			},
			validationRule: "minLen",
			wantErr:        true,
		},
		{
			name: "Trim todo title",
			todo: &model.TaskModel{
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

func TestFindByIdTask(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	todo := &model.TaskCreateRequest{
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

func TestFindByTitleOneMoreTasks(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	todos := seedTasks(t, todoService, 10)

	todoFromList := todos[:2]

	foundTasks, err := todoService.GetByTitle("Task")

	assert.Equal(t, foundTasks, todoFromList)
	assert.Len(t, foundTasks, 2)

	require.NoError(t, err)

}

func TestFindByTitleTasks(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	seedTasks := seedTasks(t, todoService, 10)

	currentTask := seedTasks[5]

	foundTasks, err := todoService.GetByTitle("Task title #5")

	assert.Equal(t, foundTasks[0], currentTask)

	require.NoError(t, err)

}

func TestFindByTitleTasksWithEmptyResult(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	seedTasks(t, todoService, 10)

	foundTasks, err := todoService.GetByTitle("Some test random text qwerty")

	assert.Len(t, foundTasks, 0)

	require.NoError(t, err)

}

// 21asdasd

func TestGetAll(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	seededTasks := seedTasks(t, todoService, 10)

	foundTasks, err := todoService.GetAll()

	assert.Len(t, foundTasks, 10)

	for index, seededTask := range seededTasks {
		assert.Equal(t, seededTask, *foundTasks[index])
	}

	require.NoError(t, err)

}

func TestDelete(t *testing.T) {
	db := test_utils.SetupTestDB(t)
	todoService := NewTaskService(db)

	seedTasks(t, todoService, 3)

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
	todoService := NewTaskService(db)

	seedTasks(t, todoService, 3)

	todoListBeforeRemove, err := todoService.GetAll()
	require.NoError(t, err)

	errOnDelete := todoService.Delete(123)
	require.NoError(t, errOnDelete)

	assert.Equal(t, errOnDelete, nil)

	todoListAfterRemove, err := todoService.GetAll()
	require.NoError(t, err)

	assert.Len(t, todoListBeforeRemove, len(todoListAfterRemove))

}
