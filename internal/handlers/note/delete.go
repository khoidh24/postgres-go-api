package note

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteNotes(c *fiber.Ctx) error {
	db := database.ConnectDB()

	var payload struct {
		IDs []string `json:"ids"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	if len(payload.IDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No IDs provided",
		})
	}

	var foundNotes []models.Note
	if err := db.Where("id IN ?", payload.IDs).Find(&foundNotes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch notes",
		})
	}

	if len(foundNotes) != len(payload.IDs) {
		// Find not exist IDs
		foundIDs := make(map[string]bool)
		for _, note := range foundNotes {
			foundIDs[note.ID] = true
		}

		var missingIDs []string
		for _, id := range payload.IDs {
			if !foundIDs[id] {
				missingIDs = append(missingIDs, id)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":    "Some notes not found",
			"missingIDs": missingIDs,
		})
	}

	// If all IDs exist, delete them
	if err := db.Where("id IN ?", payload.IDs).Delete(&models.Note{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete notes",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Notes deleted successfully",
		"count":   len(payload.IDs),
	})
}
