package repository

import (
	model "example/todo-api/internal/models"

	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

func (br *BoardRepository) FindByID(id uint) (*model.BoardModel, error) {
	var board model.BoardModel
	err := br.db.Find(&board, id).Error
	return &board, err
}

func (br *BoardRepository) FindAll() ([]model.BoardModel, error) {
	var boards []model.BoardModel
	err := br.db.Find(&boards).Error
	return boards, err
}

func (br *BoardRepository) Create(board *model.BoardModel) (*model.BoardModel, error) {
	err := br.db.Create(board).Error
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (br *BoardRepository) Delete(id uint) error {
	err := br.db.Delete(&model.BoardModel{}, id).Error
	return err
}

func (br *BoardRepository) Update(board *model.BoardModel) (*model.BoardModel, error) {
	err := br.db.Save(board).Error
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (br *BoardRepository) FindByTitle(title string) ([]model.BoardModel, error) {
	var boards []model.BoardModel
	err := br.db.Where("title LIKE ?", "%"+title+"%").Limit(2).Find(&boards).Error
	return boards, err
}
