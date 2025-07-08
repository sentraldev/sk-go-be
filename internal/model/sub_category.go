package model

import "github.com/google/uuid"

type SubCategory struct {
	Base
	CategoryUUID uuid.UUID `json:"category_uuid"`
	Category     Category  `gorm:"foreignKey:CategoryUUID" json:"category"`
	Name         string    `json:"name"`
}
