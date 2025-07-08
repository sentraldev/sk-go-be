package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	UUID      uuid.UUID `gorm:"primaryKey" json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.UUID = uuid.New()
	return nil
}
