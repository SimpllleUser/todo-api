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
	UserID uint
}

func NewTaskService(db *gorm.DB, userID uint) *TaskService {
	return &TaskService{db: db, UserID: userID}
}

func (t *TaskService) Create(req *model.TaskCreateRequest) (*model.TaskModel, error) {
	if req.BoardID == 0 {
		return nil, fmt.Errorf("board_id is required")
	}

	// TODO check if board exists
	// TODO add check if board exists and belongs to user

	task := &model.TaskModel{
		Title:       req.Title,
		Description: req.Description,
		BoardID:     req.BoardID,
		UserID:      t.UserID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}

	if err := httpUtil.Create(t.db, task); err != nil {
		return nil, err
	}

	return task, nil
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

func (t *TaskService) GetBoardTasks(boardID uint) ([]model.TaskModel, error) {
	var tasks []model.TaskModel
	err := t.db.Where("board_id = ?", boardID).Find(&tasks).Error
	return tasks, err
}
