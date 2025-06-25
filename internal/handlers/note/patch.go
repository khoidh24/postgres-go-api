package note

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ToggleActiveStatus(c *fiber.Ctx) error {
	db := database.ConnectDB()
	id := c.Params("id")

	type Body struct {
		IsActive *bool `json:"is_active"`
	}
	var body Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var note models.Note
	if err := db.First(&note, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	note.IsActive = *body.IsActive
	if err := db.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Change status success",
	})
}
