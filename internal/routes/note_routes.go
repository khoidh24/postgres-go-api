package routes

import (
	Note "postgres-go-api/internal/handlers/note"
	"postgres-go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(r fiber.Router) {
	noteRoutes := r.Group("/notes", middleware.CheckAuth)

	noteRoutes.Post("/", Note.CreateNote)
	noteRoutes.Get("/", Note.GetAllNote)
	noteRoutes.Get("/:id", Note.GetNoteById)
	noteRoutes.Put("/:id", Note.UpdateNote)
	noteRoutes.Patch("/:id", Note.ToggleActiveStatus)
	noteRoutes.Delete("/", Note.DeleteNotes)
}
