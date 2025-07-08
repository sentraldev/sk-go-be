package model

type Category struct {
	Base
	Name          string         `json:"name"`
	SubCategories *[]SubCategory `gorm:"foreignKey:CategoryUUID" json:"sub_categories"`
}
