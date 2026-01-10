package service

import (
	httpUtil "example/todo-api/internal/http_util"
	model "example/todo-api/internal/models"
	"fmt"

	"github.com/gookit/validate"
	"gorm.io/gorm"
)

type TodoService struct {
	db     *gorm.DB
	UserId uint
}

func NewTodoService(db *gorm.DB, userId uint) *TodoService {
	return &TodoService{db: db, UserId: userId}
}

func (t *TodoService) Create(todo *model.TodoCreateRequest) (*model.TodoModel, error) {
	var todoModel = &model.TodoModel{
		TodoBaseFields: todo.TodoBaseFields,
		UserId:         t.UserId,
	}
	v := validate.Struct(todo)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}
	t.db.Create(todoModel)
	return todoModel, nil
}

func (t *TodoService) Update(todo *model.TodoModel) error {
	v := validate.Struct(todo)
	if !v.Validate() {
		return fmt.Errorf("validation failed: %v", v.Errors)
	}
	return t.db.Save(todo).Error
}

func (t *TodoService) GetById(id uint) (*model.TodoModel, error) {
	return httpUtil.FindByID[model.TodoModel](t.db, id)
}

func (t *TodoService) GetByTitle(title string) ([]model.TodoModel, error) {
	var todos []model.TodoModel
	err := t.db.Where("title LIKE ?", "%"+title+"%").Limit(2).Find(&todos).Error
	return todos, err
}

func (t *TodoService) GetAll() ([]*model.TodoModel, error) {

	return httpUtil.FindAll[*model.TodoModel](t.db)
}

func (t *TodoService) Delete(id uint) error {
	return httpUtil.Delete[model.TodoModel](t.db, id)
}
