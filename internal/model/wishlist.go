package model

import "github.com/google/uuid"

type Wishlist struct {
	Base
	UserUUID    uuid.UUID `gorm:"primaryKey"`
	ProductUUID uuid.UUID `gorm:"primaryKey"`
}
