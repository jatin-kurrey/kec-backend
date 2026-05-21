package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetLeadership(c *fiber.Ctx) error {
	var leadership []models.Leadership
	if err := config.DB.Order("priority asc, created_at asc").Find(&leadership).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch leadership data"})
	}
	return c.JSON(leadership)
}

func GetAllLeadershipAdmin(c *fiber.Ctx) error {
	var leadership []models.Leadership
	if err := config.DB.Order("priority asc, created_at asc").Find(&leadership).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch leadership data"})
	}
	return c.JSON(leadership)
}

func CreateLeadership(c *fiber.Ctx) error {
	member := new(models.Leadership)
	if err := c.BodyParser(member); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Create(member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create leadership member"})
	}

	return c.Status(201).JSON(member)
}

func UpdateLeadership(c *fiber.Ctx) error {
	id := c.Params("id")
	member := new(models.Leadership)
	if err := c.BodyParser(member); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Model(&models.Leadership{}).Where("id = ?", id).Updates(member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update leadership member"})
	}

	return c.JSON(fiber.Map{"message": "Leadership member updated successfully"})
}

func DeleteLeadership(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Leadership{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete leadership member"})
	}
	return c.JSON(fiber.Map{"message": "Leadership member deleted successfully"})
}
