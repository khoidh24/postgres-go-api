package profile

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	db := database.ConnectDB()
	userID := c.Locals("userID").(string)

	var profile models.Profile
	if err := db.First(&profile, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Profile not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "success",
		"metadata": profile,
	})
}
