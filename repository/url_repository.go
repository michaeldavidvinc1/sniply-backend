package repository

import (
	"backend/models"
	"gorm.io/gorm"
)

type UrlRepository interface {
	Save(url *models.Url) (*models.Url, error)
	FindById(id string) (*models.Url, error)
	FindByShortCode(shortCode string) (*models.Url, error)
}

type urlRepository struct {
	db *gorm.DB
}

func (u urlRepository) Save(url *models.Url) (*models.Url, error) {
	err := u.db.Create(&url).Error
	return url, err
}

func (u urlRepository) FindById(id string) (*models.Url, error) {
	err := u.db.Where("id = ?", id).First(&models.Url{}).Error
	return &models.Url{}, err
}

func (u urlRepository) FindByShortCode(shortCode string) (*models.Url, error) {
	var url models.Url
	err := u.db.Where("short_code = ?", shortCode).First(&url).Error
	return &url, err
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{db}
}
