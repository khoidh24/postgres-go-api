package note

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UpdateNote(c *fiber.Ctx) error {
	db := database.ConnectDB()
	id := c.Params("id")

	type UpdateNoteInput struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CoverImage string `json:"coverImage"`
	}

	var note models.Note

	if err := db.First(&note, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	var input UpdateNoteInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	note.Title = input.Title
	note.Content = input.Content
	note.CoverImage = input.CoverImage
	note.UpdatedAt = time.Now()

	if err := db.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Modify note success",
	})
}
