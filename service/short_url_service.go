package service

import (
	"backend/helper"
	"backend/models"
	"backend/repository"
	"backend/validation"
	"github.com/google/uuid"
)

type UrlService interface {
	CreateUrl(originalUrl string) (*models.Url, error)
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
