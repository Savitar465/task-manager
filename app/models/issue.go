package models

import "time"

type Issue struct {
	ID          int       `gorm:"column:issue_id; primary_key"`
	TypeId      string    `gorm:"column:types_type_id;"`
	Title       string    `gorm:"column:title;"`
	Description string    `gorm:"column:description;"`
	StartDate   time.Time `gorm:"column:start_date;"`
	DueDate     time.Time `gorm:"column:due_date;"`
	Status      string    `gorm:"column:_status;"`
	Assignee    string    `gorm:"column:assigned_user_id;"`
	BaseModel
}
