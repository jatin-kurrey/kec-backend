package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Public
func GetPressMedia(c *fiber.Ctx) error {
	var media []models.PressMedia
	config.DB.Order("date desc").Find(&media)
	// Map virtual fields
	for i := range media {
		media[i].Image = media[i].ImageURL
	}
	return c.JSON(fiber.Map{"data": media})
}

// Admin
func CreatePressMedia(c *fiber.Ctx) error {
	media := new(models.PressMedia)
	if err := c.BodyParser(media); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Create(&media)
	return c.JSON(media)
}

func UpdatePressMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	var media models.PressMedia
	if err := config.DB.First(&media, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Media entry not found"})
	}
	if err := c.BodyParser(&media); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Save(&media)
	return c.JSON(media)
}

func DeletePressMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.PressMedia{}, "id = ?", id)
	return c.SendStatus(204)
}
