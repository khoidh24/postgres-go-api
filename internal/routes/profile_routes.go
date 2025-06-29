package routes

import (
	"postgres-go-api/internal/handlers/profile"
	"postgres-go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoutes(app *fiber.App) {
	r := app.Group("api/v1/profile", middleware.CheckAuth)

	r.Get("/profile", profile.GetProfile)
	r.Put("/profile", profile.EditProfile)
}
