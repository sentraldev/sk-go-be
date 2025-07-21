package model

import (
	"github.com/google/uuid"
)

type User struct {
	Base
	UUID           uuid.UUID `gorm:"uniqueIndex;not null" json:"uuid"`
	ExternalUserID string    `gorm:"uniqueIndex;not null" json:"external_user_id"`
	Name           string    `gorm:"not null" json:"name"`
	Phone          string    `gorm:"not null" json:"phone"`
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	Role           string    `gorm:"not null" json:"role"`
}
