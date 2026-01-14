package service

import (
	model "example/todo-api/internal/models"
	"example/todo-api/internal/repository"
	"fmt"

	"github.com/gookit/validate"
)

type TaskService struct {
	TaskRepository  *repository.TaskRepository
	UserRepository  *repository.UserRepository
	BoardRepository *repository.BoardRepository
}

func NewTaskService(taskRepository *repository.TaskRepository, userRepository *repository.UserRepository, boardRepository *repository.BoardRepository) *TaskService {
	return &TaskService{TaskRepository: taskRepository, UserRepository: userRepository, BoardRepository: boardRepository}
}

func (t *TaskService) Create(req *model.TaskCreateRequest, userId uint) (*model.TaskModel, error) {
	if req.BoardID == 0 {
		return nil, fmt.Errorf("board_id is required")
	}

	board, err := t.BoardRepository.FindByID(req.BoardID, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to find board: %v", err)
	}
	if board == nil || board.ID == 0 {
		return nil, fmt.Errorf("board with id %d does not exist", req.BoardID)
	}

	user, err := t.UserRepository.FindByID(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	if user == nil || user.ID == 0 {
		return nil, fmt.Errorf("user with id %d does not exist", userId)
	}

	task := &model.TaskModel{
		Title:       req.Title,
		Description: req.Description,
		BoardID:     req.BoardID,
		UserID:      userId,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}

	createdTask, err := t.TaskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (t *TaskService) Update(task *model.TaskModel) (*model.TaskModel, error) {
	v := validate.Struct(task)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}
	task, err := t.TaskRepository.Update(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskService) GetById(id uint, userId uint) (*model.TaskModel, error) {

	return t.TaskRepository.FindByID(id, userId)
}

func (t *TaskService) GetByTitle(title string, userId uint) ([]model.TaskModel, error) {
	var tasks []model.TaskModel

	tasks, err := t.TaskRepository.FindByTitle(title, userId)
	return tasks, err
}

func (t *TaskService) GetAll(userId uint) ([]*model.TaskModel, error) {

	tasks, err := t.TaskRepository.FindAll(userId)
	if err != nil {
		return nil, err
	}

	ptrTasks := make([]*model.TaskModel, len(tasks))
	for i := range tasks {
		ptrTasks[i] = &tasks[i]
	}
	return ptrTasks, nil
}

func (t *TaskService) Delete(id uint) error {
	return t.TaskRepository.Delete(id)
}

func (t *TaskService) GetBoardTasks(boardID uint, userId uint) ([]model.TaskModel, error) {
	return t.TaskRepository.FindByBoard(boardID, userId)
}
