package httpUtil

import "gorm.io/gorm"

func Create[T any](db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func Update[T any](db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func FindByID[T any](db *gorm.DB, id uint) (*T, error) {
	var entity T
	err := db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func FindAll[T any](db *gorm.DB) ([]T, error) {
	var entities []T
	err := db.Find(&entities).Error
	return entities, err
}

func Delete[T any](db *gorm.DB, id uint) error {
	return db.Delete(new(T), id).Error
}

func ForUser[T any](db *gorm.DB, userId uint) *gorm.DB {
	return db.Where("user_id = ?", userId)
}
