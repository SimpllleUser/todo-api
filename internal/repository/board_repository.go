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

func (br *BoardRepository) FindByID(id uint, ownerId uint) (*model.BoardModel, error) {
	var board model.BoardModel
	err := br.db.Where("id = ? AND owner_id = ?", id, ownerId).Find(&board).Error
	return &board, err
}

func (br *BoardRepository) FindAll(ownerId uint) ([]model.BoardModel, error) {
	var boards []model.BoardModel
	err := br.db.Where("owner_id = ?", ownerId).Find(&boards).Error
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

func (br *BoardRepository) FindByTitle(title string, ownerId uint) ([]model.BoardModel, error) {
	var boards []model.BoardModel
	err := br.db.Where("title LIKE ? AND owner_id = ?", "%"+title+"%", ownerId).Limit(2).Find(&boards).Error
	return boards, err
}

func (br *BoardRepository) AddUsers(boardID uint, users []model.UserModel) error {
	var board model.BoardModel
	if err := br.db.Find(&board, boardID).Error; err != nil {
		return err
	}

	if err := br.db.Model(&board).Association("Users").Append(&users); err != nil {
		return err
	}
	return nil
}
