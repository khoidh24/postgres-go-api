package note

import (
	"postgres-go-api/internal/database"
	"postgres-go-api/internal/models"
	"postgres-go-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllNote(c *fiber.Ctx) error {
	db := database.ConnectDB()

	type Note struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	var notes []Note
	var total int64

	page, limit := utils.Pagination(c)
	title := c.Query("title")
	isActive := c.QueryBool("isActive", true)
	getAll := c.QueryBool("getAll", false)

	query := db.Select("id", "title").Model(&models.Note{})

	if !getAll {
		query = query.Where("is_active = ?", isActive)
	}

	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	query.Count(&total)

	query.Offset((page - 1) * limit).Limit(limit).Find(&notes)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "success",
		"metadata": notes,
		"page":     page,
		"limit":    limit,
		"total":    total,
	})
}

func GetNoteById(c *fiber.Ctx) error {
	db := database.ConnectDB()
	id := c.Params("id")

	var note models.Note

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    note,
	})
}
