package model

type TodoModel struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Title       string `gorm:"not null" json:"title" example:"Buy groceries"`
	Description string `gorm:"null" json:"description" example:"Milk, eggs, bread"`
	Completed   bool   `gorm:"default:false" json:"completed" example:"false"`
}

func (TodoModel) TableName() string {
	return "todos"
}
