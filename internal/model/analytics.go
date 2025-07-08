package model

type Analytics struct {
	Base
	MetricID uint   `gorm:"primaryKey"`
	Type     string `gorm:"not null"`
	Value    float64
	Date     string `gorm:"not null"`
}
