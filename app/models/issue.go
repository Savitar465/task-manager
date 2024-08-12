package models

type Issue struct {
	ID          int    `gorm:"column:issue_id; primary_key"`
	TypeId      string `gorm:"column:types_type_id;"`
	Title       string `gorm:"column:title;"`
	Description string `gorm:"column:description;"`
	StartDate   string `gorm:"column:start_date;"`
	DueDate     string `gorm:"column:due_date;"`
	Status      string `gorm:"column:_status;"`
	Assignee    string `gorm:"column:assigned_user_id;"`
	StageId     string `gorm:"column:stages_stage_id;"`
	BoardId     string `gorm:"column:boards_board_id;"`
	BaseModel
}
