package main

import (
	"backend/config"
	"backend/controller"
	"backend/middleware"
	"backend/models"
	"backend/routes"
	"backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Url{})

	r := gin.Default()

	// Middlewares
	r.Use(middleware.LoggerMiddleware())

	urlService := service.NewUrlService(db)
	urlController := controller.NewUrlController(urlService)

	// Routes
	routes.Routes(r, urlController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
