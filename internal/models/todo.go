package model

type TodoModel struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null"`
	Description string `gorm:"null"`
	Completed   bool   `gorm:"default:false"`
}

func (TodoModel) TableName() string {
	return "todos"
}
