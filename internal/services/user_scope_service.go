package service

import (
	"gorm.io/gorm"
)

type UserScopeService struct {
	db *gorm.DB
}

func NewUserScopeService(db *gorm.DB) *UserScopeService {
	return &UserScopeService{db: db}
}

func (s *UserScopeService) ForUser(userID uint) *UserScope {
	return &UserScope{
		db:     s.db.Where("user_id = ?", userID),
		userID: userID,
	}
}

func (s *UserScopeService) ForUserOwner(userID uint) *UserScopeBoard {
	return &UserScopeBoard{
		db:      s.db.Where("owner_id = ?", userID),
		OwnerID: userID,
	}
}

type UserScope struct {
	db     *gorm.DB
	userID uint
}

type UserScopeBoard struct {
	db      *gorm.DB
	OwnerID uint
}

func (u *UserScope) Task() *TaskService {
	return &TaskService{db: u.db, UserID: u.userID}
}

func (u *UserScopeBoard) Board() *BoardService {
	return &BoardService{db: u.db, OwnerID: u.OwnerID}
}
