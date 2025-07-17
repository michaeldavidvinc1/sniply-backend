package service

import (
	"backend/helper"
	"backend/models"
	"backend/repository"
	"backend/validation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrlService interface {
	CreateUrl(req validation.CreateUrlRequest) (*models.Url, error)
	GetOriginalUrl(shortCode string) (*models.Url, error)
}

type urlService struct {
	urlRepo repository.UrlRepository
}

func (u *urlService) CreateUrl(req validation.CreateUrlRequest) (*models.Url, error) {
	shortCode := helper.GenerateShortCode()
	newUrl := &models.Url{
		ID:          uuid.NewString(),
		OriginalURL: req.OriginalUrl,
		ShortCode:   shortCode,
	}
	return u.urlRepo.Save(newUrl)
}

func (u *urlService) GetOriginalUrl(shortCode string) (*models.Url, error) {
	result, err := u.urlRepo.FindByShortCode(shortCode)

	return result, err
}

func NewUrlService(db *gorm.DB) UrlService {
	return &urlService{urlRepo: repository.NewUrlRepository(db)}

}
