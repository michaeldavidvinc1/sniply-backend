package models

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	ID          string  `gorm:"primaryKey"`
	OriginalURL string  `gorm:"type:text; not null"`
	ShortCode   string  `gorm:"uniqueIndex; not null"`
	UserID      *string `gorm:"null"`
	CreatedAt   time.Time
	ExpiresAt   time.Time `gorm:"null"`
}
