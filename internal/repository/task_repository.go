package repository

import (
	model "example/todo-api/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) FindByID(id uint, userId uint) (*model.TaskModel, error) {
	var task model.TaskModel
	err := tr.db.Where("id = ? AND user_id = ?", id, userId).Find(&task).Error
	return &task, err
}

func (tr *TaskRepository) FindAll(userId uint) ([]model.TaskModel, error) {
	var tasks []model.TaskModel
	err := tr.db.Where("user_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (tr *TaskRepository) Create(task *model.TaskModel) (*model.TaskModel, error) {
	err := tr.db.Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *TaskRepository) Delete(id uint) error {
	err := tr.db.Delete(&model.TaskModel{}, id).Error
	return err
}

func (tr *TaskRepository) Update(task *model.TaskModel) (*model.TaskModel, error) {
	err := tr.db.Save(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *TaskRepository) FindByTitle(title string, userId uint) ([]model.TaskModel, error) {
	var tasks []model.TaskModel
	err := tr.db.Where("title LIKE ? AND user_id = ?", "%"+title+"%", userId).Limit(2).Find(&tasks).Error
	return tasks, err
}

func (tr *TaskRepository) FindByBoard(boardID uint, userId uint) ([]model.TaskModel, error) {
	var tasks []model.TaskModel
	err := tr.db.Where("board_id = ? AND user_id = ?", boardID, userId).Find(&tasks).Error
	return tasks, err
}
