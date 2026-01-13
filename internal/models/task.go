package model

type TaskBaseFields struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Title       string `gorm:"not null" json:"title" validate:"required|minLen:3" filter:"trim" message:"required:{field} is required" label:"Title" example:"Buy groceries"`
	Description string `gorm:"null" json:"description" example:"Milk, eggs, bread"`
	Completed   bool   `gorm:"default:false" json:"completed" example:"false"`
	TaskBoardId
}

type TaskBoardId struct {
	BoardID uint `json:"board_id" gorm:"not null"`
}

type TaskModel struct {
	TaskBaseFields
	UserId uint `gorm:"not null" json:"user_id" example:"1"`
}
type TaskCreateRequest struct {
	TaskBaseFields
}

func (TaskModel) TableName() string {
	return "tasks"
}
