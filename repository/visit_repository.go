package repository

import (
	"backend/models"
	"gorm.io/gorm"
)

type VisitRepository interface {
	Save(visit *models.Visit) (*models.Visit, error)
}

type visitRepository struct {
	db *gorm.DB
}

func (v visitRepository) Save(visit *models.Visit) (*models.Visit, error) {
	err := v.db.Save(visit).Error
	return visit, err
}
