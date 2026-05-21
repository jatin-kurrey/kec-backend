package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Public
func GetAlumni(c *fiber.Ctx) error {
	var alumni []models.Alumni
	config.DB.Order("created_at desc").Find(&alumni)
	// Map virtual fields for frontend compatibility
	for i := range alumni {
		alumni[i].Photo = alumni[i].ImageURL
	}
	return c.JSON(fiber.Map{"data": alumni})
}

func GetAlumniStats(c *fiber.Ctx) error {
	var stats []models.AlumniStat
	config.DB.Order("created_at asc").Find(&stats)
	return c.JSON(fiber.Map{"data": stats})
}

// Admin
func CreateAlumni(c *fiber.Ctx) error {
	alumni := new(models.Alumni)
	if err := c.BodyParser(alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Create(&alumni)
	return c.JSON(alumni)
}

func UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var alumni models.Alumni
	if err := config.DB.First(&alumni, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni not found"})
	}
	if err := c.BodyParser(&alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Save(&alumni)
	return c.JSON(alumni)
}

func DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Alumni{}, "id = ?", id)
	return c.SendStatus(204)
}

func UpdateAlumniStats(c *fiber.Ctx) error {
	var stats []models.AlumniStat
	if err := c.BodyParser(&stats); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	// Simplified: delete all and recreate or update one by one
	// For now, let's just create if it doesn't exist
	for _, s := range stats {
		if s.ID == uuid.Nil {
			config.DB.Create(&s)
		} else {
			config.DB.Save(&s)
		}
	}
	return c.JSON(stats)
}
