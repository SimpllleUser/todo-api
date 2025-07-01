package model

import "gorm.io/gorm"

type TodoModel struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null"`
	Description string `gorm:"null"`
	Completed   bool   `gorm:"default:false"`
}

func (TodoModel) TableName() string {
	return "todos"
}

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (t *TodoService) Create(todo *TodoModel) error {
	return t.db.Create(todo).Error
}

func (t *TodoService) GetById(id uint) (*TodoModel, error) {
	var todo TodoModel
	err := t.db.Find(&todo, id).Error
	return &todo, err
}

func (t *TodoService) GetByTitle(title string) (*[]TodoModel, error) {
	var todos []TodoModel
	err := t.db.Where("title LIKE ?", "%"+title+"%").Limit(2).Find(&todos).Error
	return &todos, err
}

func (t *TodoService) GetAll() ([]TodoModel, error) {
	var todos = []TodoModel{}

	err := t.db.Find(&todos).Error
	return todos, err
}

func (t *TodoService) Update(todo *TodoModel) error {
	return t.db.Save(todo).Error
}

func (t *TodoService) Delete(id uint) error {
	return t.db.Delete(&TodoModel{}, id).Error
}
