package service

import (
	httpUtil "example/todo-api/internal/http_util"
	model "example/todo-api/internal/models"
	"fmt"

	"github.com/gookit/validate"
	"gorm.io/gorm"
)

type TaskService struct {
	db     *gorm.DB
	UserId uint
}

func NewTaskService(db *gorm.DB, userId uint) *TaskService {
	return &TaskService{db: db, UserId: userId}
}

func (t *TaskService) Create(task *model.TaskCreateRequest) (*model.TaskModel, error) {
	var taskModel = &model.TaskModel{
		TaskBaseFields: task.TaskBaseFields,
		UserId:         t.UserId,
	}
	v := validate.Struct(task)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}
	httpUtil.Create[model.TaskModel](t.db, taskModel)
	return taskModel, nil
}

func (t *TaskService) Update(task *model.TaskModel) error {
	v := validate.Struct(task)
	if !v.Validate() {
		return fmt.Errorf("validation failed: %v", v.Errors)
	}
	return t.db.Save(task).Error
}

func (t *TaskService) GetById(id uint) (*model.TaskModel, error) {
	return httpUtil.FindByID[model.TaskModel](t.db, id)
}

func (t *TaskService) GetByTitle(title string) ([]model.TaskModel, error) {
	var tasks []model.TaskModel
	err := t.db.Where("title LIKE ?", "%"+title+"%").Limit(2).Find(&tasks).Error
	return tasks, err
}

func (t *TaskService) GetAll() ([]*model.TaskModel, error) {

	return httpUtil.FindAll[*model.TaskModel](t.db)
}

func (t *TaskService) Delete(id uint) error {
	return httpUtil.Delete[model.TaskModel](t.db, id)
}
