package profile

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func EditProfile(c *fiber.Ctx) error {

	db := database.ConnectDB()
	userID := c.Locals("userID").(string)

	var input struct {
		FullName string `json:"full_name"`
		Bio      string `json:"bio"`
		CoverUrl string `json:"cover_url"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var profile models.Profile
	if err := db.FirstOrCreate(&profile, models.Profile{ID: userID}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	profile.FullName = input.FullName
	profile.Bio = input.Bio
	profile.CoverUrl = input.CoverUrl

	db.Save(&profile)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
