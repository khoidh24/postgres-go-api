package note

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CreateNoteInput struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CoverImage string `json:"coverImage"`
}

func CreateNote(c *fiber.Ctx) error {
	db := database.ConnectDB()
	var input CreateNoteInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	note := models.Note{
		Title:      input.Title,
		Content:    input.Content,
		CoverImage: input.CoverImage,
	}

	if err := db.Create(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    note,
	})
}
