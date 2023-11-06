package model

import (
	"time"

	"gorm.io/gorm"
)

type Redirect struct {
	gorm.Model

	ShortURL    string `gorm:"uniqueIndex;not null"`
	OriginalURL string `gorm:"not null"`
	Hits        uint   `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
