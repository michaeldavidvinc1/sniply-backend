package routes

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, urlController *controller.ShortUrlController) {
	api := router.Group("/api")
	{
		api.POST("/shorten", urlController.CreateUrl)
	}

	router.GET("/:shortCode", urlController.RedirectUrl)
}
