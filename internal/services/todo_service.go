package service

import (
	model "example/todo-api/internal/models"

	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (t *TodoService) Create(todo *model.TodoModel) error {
	return t.db.Create(todo).Error
}

func (t *TodoService) GetById(id uint) (*model.TodoModel, error) {
	var todo *model.TodoModel
	err := t.db.Find(&todo, id).Error
	return todo, err
}

func (t *TodoService) GetByTitle(title string) (*[]*model.TodoModel, error) {
	var todos []*model.TodoModel
	err := t.db.Where("title LIKE ?", "%"+title+"%").Limit(2).Find(&todos).Error
	return &todos, err
}

func (t *TodoService) GetAll() ([]*model.TodoModel, error) {
	var todos = []*model.TodoModel{}

	err := t.db.Find(&todos).Error
	return todos, err
}

func (t *TodoService) Update(todo *model.TodoModel) error {
	return t.db.Save(todo).Error
}

func (t *TodoService) Delete(id uint) error {
	return t.db.Delete(&model.TodoModel{}, id).Error
}
