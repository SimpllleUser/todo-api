package model

import "gorm.io/gorm"

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (t *TodoService) Create(todo *TodoModel) error {
	return t.db.Create(todo).Error
}

func (t *TodoService) GetAll() ([]TodoModel, error) {
	var todos = []TodoModel{}

	err := t.db.Find(&todos).Error
	return todos, err
}
