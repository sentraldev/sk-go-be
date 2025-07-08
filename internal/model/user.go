package model

type User struct {
	Base
	Email        string `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null" json:"role"`
}
