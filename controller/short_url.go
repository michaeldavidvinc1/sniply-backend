package controller

import (
	"backend/service"
	"backend/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ShortUrlController struct {
	UrlService service.UrlService
	Validator  *validator.Validate
}

func NewUrlController(service service.UrlService) *ShortUrlController {
	return &ShortUrlController{
		UrlService: service,
		Validator:  validator.New(),
	}
}

func (uc *ShortUrlController) CreateUrl(c *gin.Context) {
	var req validation.CreateUrlRequest

	// Bind dan cek error format JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	// Validasi manual
	if err := uc.Validator.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var messages []string
		for _, ve := range validationErrors {
			messages = append(messages, ve.Field()+" is invalid ("+ve.Tag()+")")
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": messages})
		return
	}

	// Call service
	result, err := uc.UrlService.CreateUrl(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
