package routes

import (
	"postgres-go-api/internal/handlers/profile"
	"postgres-go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoutes(r fiber.Router) {
	profileRoutes := r.Group("/profile", middleware.CheckAuth)

	profileRoutes.Get("/", profile.GetProfile)
	profileRoutes.Put("/", profile.EditProfile)
}
