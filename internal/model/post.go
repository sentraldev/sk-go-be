package model

import "time"

type Post struct {
	Base
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Author  string `gorm:"not null"`
	Date    time.Time
	Tags    []string `gorm:"type:text[]"`
}
