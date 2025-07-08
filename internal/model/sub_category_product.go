package model

import "github.com/google/uuid"

type SubCategoryProduct struct {
	Base
	ProductUUID     uuid.UUID `json:"product_uuid"`
	SubCategoryUUID uuid.UUID `json:"sub_category_uuid"`
}
