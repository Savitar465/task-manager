package models

import "time"

type BaseModel struct {
	CreatedBy string    `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedBy string    `gorm:"column:updated_by;" json:"updatedBy"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updatedAt"`
}
