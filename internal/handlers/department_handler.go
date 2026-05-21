package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetDepartments returns all active departments
func GetDepartments(c *fiber.Ctx) error {
	var departments []models.Department
	if err := config.DB.Where("is_active = ?", true).Order("name asc").Find(&departments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch departments"})
	}
	return c.JSON(departments)
}

// GetDepartmentByShortName returns a specific department by its short name
func GetDepartmentByShortName(c *fiber.Ctx) error {
	shortName := c.Params("shortName")
	var dept models.Department
	if err := config.DB.Where("LOWER(short_name) = LOWER(?)", shortName).First(&dept).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Department not found"})
	}
	return c.JSON(dept)
}

// GetAllDepartmentsAdmin returns all departments for admin management
func GetAllDepartmentsAdmin(c *fiber.Ctx) error {
	var departments []models.Department
	if err := config.DB.Order("name asc").Find(&departments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch departments"})
	}
	return c.JSON(departments)
}

// CreateDepartment creates a new department
func CreateDepartment(c *fiber.Ctx) error {
	dept := new(models.Department)
	if err := c.BodyParser(dept); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Create(&dept).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create department"})
	}

	return c.Status(201).JSON(dept)
}

// UpdateDepartment updates an existing department
func UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	dept := new(models.Department)
	if err := c.BodyParser(dept); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Model(&models.Department{}).Where("id = ?", id).Updates(dept).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update department"})
	}

	return c.JSON(fiber.Map{"message": "Department updated successfully"})
}

// DeleteDepartment deletes a department
func DeleteDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Department{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete department"})
	}
	return c.JSON(fiber.Map{"message": "Department deleted successfully"})
}
