package model

type Cart struct {
	Base
	UserID    uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  int  `gorm:"not null"`
}
