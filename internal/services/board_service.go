package service

import (
	model "example/todo-api/internal/models"
	"example/todo-api/internal/repository"
	"fmt"

	"github.com/gookit/validate"
)

type BoardService struct {
	UserRepository  *repository.UserRepository
	BoardRepository *repository.BoardRepository
}

func NewBoardService(userRepository *repository.UserRepository, boardRepository *repository.BoardRepository) *BoardService {
	return &BoardService{
		UserRepository:  userRepository,
		BoardRepository: boardRepository,
	}
}

func (t *BoardService) Create(req *model.BoardCreateRequest, ownerId uint) (*model.BoardModel, error) {
	user, err := t.UserRepository.FindByID(ownerId)
	if err != nil {
		return nil, fmt.Errorf("failed to find owner: %v", err)
	}
	if user == nil || user.ID == 0 {
		return nil, fmt.Errorf("owner with id %d does not exist", ownerId)
	}

	board := &model.BoardModel{
		Title:   req.Title,
		OwnerID: ownerId,
	}

	v := validate.Struct(board)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}

	createdBoard, err := t.BoardRepository.Create(board)
	if err != nil {
		return nil, err
	}

	return createdBoard, nil
}

func (t *BoardService) Update(req *model.BoardModel) (*model.BoardModel, error) {
	v := validate.Struct(req)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}
	board, err := t.BoardRepository.Update(req)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (t *BoardService) GetById(id uint, ownerId uint) (*model.BoardModel, error) {
	return t.BoardRepository.FindByID(id, ownerId)
}

func (t *BoardService) GetByTitle(title string, ownerId uint) ([]model.BoardModel, error) {
	return t.BoardRepository.FindByTitle(title, ownerId)
}

func (t *BoardService) GetAll(ownerId uint) ([]*model.BoardModel, error) {

	boards, err := t.BoardRepository.FindAll(ownerId)
	if err != nil {
		return nil, err
	}

	ptrBoards := make([]*model.BoardModel, len(boards))
	for i := range boards {
		ptrBoards[i] = &boards[i]
	}
	return ptrBoards, nil
}

func (t *BoardService) Delete(id uint) error {
	return t.BoardRepository.Delete(id)
}

func (t *BoardService) AddUsers(boardID uint, ownerId uint, userIDs []uint) (*model.BoardModel, error) {
	board, err := t.BoardRepository.FindByID(boardID, ownerId)
	if err != nil {
		return nil, err
	}
	if board == nil || board.ID == 0 {
		return nil, fmt.Errorf("board with id %d does not exist", boardID)
	}

	users := make([]model.UserModel, 0, len(userIDs))
	for _, userID := range userIDs {
		user, err := t.UserRepository.FindByID(userID)
		if err != nil {
			return nil, fmt.Errorf("failed to find user with id %d: %v", userID, err)
		}
		if user == nil || user.ID == 0 {
			return nil, fmt.Errorf("user with id %d does not exist", userID)
		}
		users = append(users, *user)
	}

	if err := t.BoardRepository.AddUsers(boardID, users); err != nil {
		return nil, err
	}

	boardWithUsers, err := t.BoardRepository.FindByIDWithUsers(boardID)
	if err != nil {
		return nil, err
	}

	return boardWithUsers, nil
}
