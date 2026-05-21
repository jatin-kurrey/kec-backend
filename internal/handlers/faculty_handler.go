package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetFaculty returns all active faculty members
func GetFaculty(c *fiber.Ctx) error {
	var faculty []models.Faculty
	if err := config.DB.Where("is_active = ?", true).Order("s_no asc, name asc").Find(&faculty).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch faculty"})
	}
	return c.JSON(faculty)
}

// GetAllFacultyAdmin returns all faculty members for admin management
func GetAllFacultyAdmin(c *fiber.Ctx) error {
	var faculty []models.Faculty
	if err := config.DB.Order("s_no asc, name asc").Find(&faculty).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch faculty"})
	}
	return c.JSON(faculty)
}

// CreateFaculty creates a new faculty member
func CreateFaculty(c *fiber.Ctx) error {
	member := new(models.Faculty)
	if err := c.BodyParser(member); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Create(member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create faculty member"})
	}

	return c.Status(201).JSON(member)
}

// UpdateFaculty updates an existing faculty member
func UpdateFaculty(c *fiber.Ctx) error {
	id := c.Params("id")
	member := new(models.Faculty)
	if err := c.BodyParser(member); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Model(&models.Faculty{}).Where("id = ?", id).Updates(member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update faculty member"})
	}

	return c.JSON(fiber.Map{"message": "Faculty member updated successfully"})
}

// DeleteFaculty deletes a faculty member
func DeleteFaculty(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Faculty{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete faculty member"})
	}
	return c.JSON(fiber.Map{"message": "Faculty member deleted successfully"})
}
