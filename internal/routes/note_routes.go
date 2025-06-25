package routes

import (
	Note "postgres-go-api/internal/handlers/note"
	"postgres-go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(app *fiber.App) {
	note := app.Group("/api/v1/notes", middleware.CheckAuth)
	note.Post("/", Note.CreateNote)
	note.Get("/", Note.GetAllNote)
	note.Get("/:id", Note.GetNoteById)
	note.Put("/:id", Note.UpdateNote)
	note.Patch("/:id", Note.ToggleActiveStatus)
	note.Delete("/", Note.DeleteNotes)
}
