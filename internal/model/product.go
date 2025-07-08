package model

import "github.com/google/uuid"

type Product struct {
	Base
	Name          string         `gorm:"not null" json:"name"`
	CategoryUUID  uuid.UUID      `gorm:"not null" json:"category_uuid"`
	Category      Category       `gorm:"foreignKey:CategoryUUID" json:"category"`
	Price         float64        `gorm:"not null" json:"price"`
	Description   string         `json:"description"`
	SubCategories *[]SubCategory `gorm:"many2many:sub_category_products" json:"sub_categories"`
	ImageURLs     []string       `gorm:"type:text[]"`
}
