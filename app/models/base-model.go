package models

type BaseModel struct {
	CreatedBy string `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt string `gorm:"column:created_at;" json:"createdAt"`
	UpdatedBy string `gorm:"column:updated_by;" json:"updatedBy"`
	UpdatedAt string `gorm:"column:updated_at;" json:"updatedAt"`
}
