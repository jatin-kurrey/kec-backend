package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Public Handlers

func GetCampusFacilities(c *fiber.Ctx) error {
	var facilities []models.CampusFacility
	category := c.Query("category")
	
	db := config.DB.Where("is_active = ?", true).Order("sort_order asc")
	if category != "" {
		db = db.Where("category = ?", category)
	}
	
	if err := db.Find(&facilities).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch facilities"})
	}
	return c.JSON(facilities)
}

func GetCampusStats(c *fiber.Ctx) error {
	var stats []models.CampusStat
	if err := config.DB.Order("sort_order asc").Find(&stats).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch stats"})
	}
	return c.JSON(stats)
}

// Admin Handlers

func CreateCampusFacility(c *fiber.Ctx) error {
	facility := new(models.CampusFacility)
	if err := c.BodyParser(facility); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Create(facility).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create facility"})
	}
	return c.Status(201).JSON(facility)
}

func UpdateCampusFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	facility := new(models.CampusFacility)
	if err := c.BodyParser(facility); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Model(&models.CampusFacility{}).Where("id = ?", id).Updates(facility).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update facility"})
	}
	return c.JSON(fiber.Map{"message": "Facility updated successfully"})
}

func DeleteCampusFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.CampusFacility{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete facility"})
	}
	return c.JSON(fiber.Map{"message": "Facility deleted successfully"})
}

func CreateCampusStat(c *fiber.Ctx) error {
	stat := new(models.CampusStat)
	if err := c.BodyParser(stat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Create(stat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create stat"})
	}
	return c.Status(201).JSON(stat)
}

func UpdateCampusStat(c *fiber.Ctx) error {
	id := c.Params("id")
	stat := new(models.CampusStat)
	if err := c.BodyParser(stat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Model(&models.CampusStat{}).Where("id = ?", id).Updates(stat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update stat"})
	}
	return c.JSON(fiber.Map{"message": "Stat updated successfully"})
}

func DeleteCampusStat(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.CampusStat{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete stat"})
	}
	return c.JSON(fiber.Map{"message": "Stat deleted successfully"})
}
