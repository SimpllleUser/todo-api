package model

type TodoBaseFields struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Title       string `gorm:"not null" json:"title" validate:"required|minLen:3" filter:"trim" message:"required:{field} is required" label:"Title" example:"Buy groceries"`
	Description string `gorm:"null" json:"description" example:"Milk, eggs, bread"`
	Completed   bool   `gorm:"default:false" json:"completed" example:"false"`
	TodoProjectId
}

type TodoProjectId struct {
	ProjectId uint `gorm:"not null" json:"project_id" example:"1"`
}

type TodoModel struct {
	TodoBaseFields
	UserId uint `gorm:"not null" json:"user_id" example:"1"`
}
type TodoCreateRequest struct {
	TodoBaseFields
}

func (TodoModel) TableName() string {
	return "todos"
}
