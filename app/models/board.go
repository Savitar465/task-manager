package models

type Board struct {
	ID     int    `gorm:"column:board_id; primary_key"`
	Name   string `gorm:"column:name;"`
	Status string `gorm:"column:_status;"`
	BaseModel
}
