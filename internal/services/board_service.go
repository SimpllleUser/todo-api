package service

import (
	httpUtil "example/todo-api/internal/http_util"
	model "example/todo-api/internal/models"
	"fmt"

	"github.com/gookit/validate"
	"gorm.io/gorm"
)

type BoardService struct {
	db         *gorm.DB
	originalDB *gorm.DB // додай це
	OwnerID    uint
}

func NewBoardService(db *gorm.DB, UserId uint) *BoardService {
	return &BoardService{
		db:         db,
		originalDB: db.Session(&gorm.Session{NewDB: true}), // новий чистий DB
		OwnerID:    UserId,
	}
}

func (t *BoardService) Create(board *model.BoardCreateRequest) (*model.BoardModel, error) {

	var BoardModel = &model.BoardModel{
		Title:   board.Title,
		OwnerID: t.OwnerID,
	}

	fmt.Println("Creating board for owner ID:", t.OwnerID)

	v := validate.Struct(board)
	if !v.Validate() {
		return nil, fmt.Errorf("validation failed: %v", v.Errors)
	}
	httpUtil.Create[model.BoardModel](t.db, BoardModel)
	return BoardModel, nil
}

func (t *BoardService) Update(board *model.BoardModel) error {
	v := validate.Struct(board)
	if !v.Validate() {
		return fmt.Errorf("validation failed: %v", v.Errors)
	}
	return t.db.Save(board).Error
}

func (t *BoardService) GetById(id uint) (*model.BoardModel, error) {
	return httpUtil.FindByID[model.BoardModel](t.db, id)
}

func (t *BoardService) GetByTitle(title string) ([]model.BoardModel, error) {
	var boards []model.BoardModel
	err := t.db.Where("title LIKE ?", "%"+title+"%").Limit(2).Find(&boards).Error
	return boards, err
}

func (t *BoardService) GetAll() ([]*model.BoardModel, error) {

	return httpUtil.FindAll[*model.BoardModel](t.db)
}

func (t *BoardService) Delete(id uint) error {
	return httpUtil.Delete[model.BoardModel](t.db, id)
}

func (t *BoardService) AddUsers(boardID uint, userIDs []uint) (*model.BoardModel, error) {
	db := t.originalDB

	var board model.BoardModel
	if err := db.First(&board, boardID).Error; err != nil {
		return nil, err
	}

	users := make([]model.UserModel, len(userIDs))
	for i, userID := range userIDs {
		users[i] = model.UserModel{ID: userID}
	}

	if err := db.Model(&board).Association("Users").Append(&users); err != nil {
		return nil, err
	}

	if err := db.Preload("Users").First(&board, boardID).Error; err != nil {
		return nil, err
	}

	return &board, nil
}
