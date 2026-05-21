package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Public
func GetPlacementStats(c *fiber.Ctx) error {
	var stats []models.PlacementStat
	config.DB.Find(&stats)
	return c.JSON(fiber.Map{"data": stats})
}

func GetRecruiters(c *fiber.Ctx) error {
	var recruiters []models.Recruiter
	config.DB.Find(&recruiters)
	// Map virtual fields
	for i := range recruiters {
		recruiters[i].Logo = recruiters[i].LogoURL
	}
	return c.JSON(fiber.Map{"data": recruiters})
}

func GetPlacementTestimonials(c *fiber.Ctx) error {
	var testimonials []models.PlacementTestimonial
	config.DB.Find(&testimonials)
	// Map virtual fields
	for i := range testimonials {
		testimonials[i].Image = testimonials[i].ImageURL
	}
	return c.JSON(fiber.Map{"data": testimonials})
}

// Admin
func UpdatePlacementStats(c *fiber.Ctx) error {
	var stats []models.PlacementStat
	if err := c.BodyParser(&stats); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	for _, s := range stats {
		if s.ID == uuid.Nil {
			config.DB.Create(&s)
		} else {
			config.DB.Save(&s)
		}
	}
	return c.JSON(stats)
}

func CreateRecruiter(c *fiber.Ctx) error {
	recruiter := new(models.Recruiter)
	if err := c.BodyParser(recruiter); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Create(&recruiter)
	return c.JSON(recruiter)
}

func UpdateRecruiter(c *fiber.Ctx) error {
	id := c.Params("id")
	var recruiter models.Recruiter
	if err := config.DB.First(&recruiter, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Recruiter not found"})
	}
	if err := c.BodyParser(&recruiter); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Save(&recruiter)
	return c.JSON(recruiter)
}

func DeleteRecruiter(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Recruiter{}, "id = ?", id)
	return c.SendStatus(204)
}

func CreatePlacementTestimonial(c *fiber.Ctx) error {
	testimonial := new(models.PlacementTestimonial)
	if err := c.BodyParser(testimonial); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Create(&testimonial)
	return c.JSON(testimonial)
}

func UpdatePlacementTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	var testimonial models.PlacementTestimonial
	if err := config.DB.First(&testimonial, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Testimonial not found"})
	}
	if err := c.BodyParser(&testimonial); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Save(&testimonial)
	return c.JSON(testimonial)
}

func DeletePlacementTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.PlacementTestimonial{}, "id = ?", id)
	return c.SendStatus(204)
}
