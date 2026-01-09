package service

import "gorm.io/gorm"

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

type UserScope struct {
	db     *gorm.DB
	userID uint
}

func (u *UserScope) Todo() *TodoService {
	return &TodoService{db: u.db, UserId: u.userID}
}
