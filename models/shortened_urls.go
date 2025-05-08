package models

import (
	"gorm.io/gorm"
	"time"
)

type ShortenedURL struct {
	gorm.Model
	ID          string  `gorm:"primaryKey"`
	OriginalURL string  `gorm:"type:text; not null"`
	ShortCode   string  `gorm:"uniqueIndex; not null"`
	UserID      *string `gorm:"null"` // Nullable (jika tanpa login)
	VisitCount  int     `gorm:"default:0"`
	CreatedAt   time.Time
	ExpiresAt   time.Time `gorm:"null"`
}
