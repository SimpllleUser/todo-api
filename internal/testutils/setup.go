package test_utils

import (
	model "example/todo-api/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.TaskModel{}, &model.UserModel{})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := db.Exec("DELETE FROM tasks").Error
		require.NoError(t, err)

		err = db.Exec("DELETE FROM users").Error
		require.NoError(t, err)
	})
	return db
}
