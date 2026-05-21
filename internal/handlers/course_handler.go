package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetCourses returns all active courses
func GetCourses(c *fiber.Ctx) error {
	var courses []models.Course
	if err := config.DB.Where("is_active = ?", true).Order("created_at asc").Find(&courses).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch courses"})
	}
	return c.JSON(courses)
}

// GetAllCoursesAdmin returns all courses for admin management
func GetAllCoursesAdmin(c *fiber.Ctx) error {
	var courses []models.Course
	if err := config.DB.Order("created_at desc").Find(&courses).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch courses"})
	}
	return c.JSON(courses)
}

// CreateCourse creates a new course
func CreateCourse(c *fiber.Ctx) error {
	course := new(models.Course)
	if err := c.BodyParser(course); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Create(course).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create course"})
	}

	return c.Status(201).JSON(course)
}

// UpdateCourse updates an existing course
func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	course := new(models.Course)
	if err := c.BodyParser(course); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := config.DB.Model(&models.Course{}).Where("id = ?", id).Updates(course).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update course"})
	}

	return c.JSON(fiber.Map{"message": "Course updated successfully"})
}

// DeleteCourse deletes a course
func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Course{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete course"})
	}
	return c.JSON(fiber.Map{"message": "Course deleted successfully"})
}
