package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterV1Routes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	NoteRoutes(v1)
	ProfileRoutes(v1)
}
