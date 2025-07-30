package model

import (
	utils "sk-go-be/internal/utils/type"

	"github.com/google/uuid"
)

type Product struct {
	Base
	Name         string           `gorm:"not null" json:"name"`
	CategoryUUID uuid.UUID        `gorm:"not null" json:"category_uuid"`
	Category     Category         `gorm:"foreignKey:CategoryUUID" json:"category"`
	Price        float64          `gorm:"not null" json:"price"`
	Description  string           `json:"description"`
	Highlight    string           `json:"highlight"`
	ImageURLs    utils.JSONBArray `gorm:"type:jsonb" json:"image_urls"`
	BrandUUID    uuid.UUID        `gorm:"not null" json:"brand_uuid"`
	Brand        Brand            `gorm:"foreignKey:BrandUUID" json:"brand"`
}
