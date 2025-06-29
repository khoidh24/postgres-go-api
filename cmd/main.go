package main

import (
	"postgres-go-api/internal/config"
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"
	"postgres-go-api/internal/routes"

	"github.com/gofiber/fiber/v2"

	"os"
)

func main() {
	config.LoadConfig()

	app := fiber.New()
	db := database.ConnectDB()
	db.AutoMigrate(&models.Note{})

	routes.RegisterV1Routes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
