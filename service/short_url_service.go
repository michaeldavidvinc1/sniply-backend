package service

import (
	"backend/helper"
	"backend/models"
	"backend/repository"
	"errors"
)

type UrlService interface {
	CreateUrl(originalUrl string) (*models.Url, error)
}

type urlService struct {
	urlRepo repository.UrlRepository
}

func NewUrlService(urlRepo repository.UrlRepository) UrlService {
	return &urlService{urlRepo}
}

func (u *urlService) CreateUrl(originalUrl string) (*models.Url, error) {
	if originalUrl == "" {
		return nil, errors.New("original URL cannot be empty")
	}
	shortCode := helper.GenerateShortCode()
	newUrl := &models.Url{
		ID:          uuid.NewString(),
		OriginalURL: originalUrl,
		ShortCode:   shortCode,
	}
	return u.urlRepo.Save(newUrl)
}
