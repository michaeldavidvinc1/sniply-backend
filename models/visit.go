package models

import (
	"gorm.io/gorm"
	"time"
)

type Visit struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	ShortenedURLID string `gorm:"index"`
	AccessedAt     time.Time
	IPAddress      string `gorm:"null"`
	Referrer       string `gorm:"null"`
}
