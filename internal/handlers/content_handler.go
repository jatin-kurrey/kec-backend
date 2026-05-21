package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Notices
func GetNotices(c *fiber.Ctx) error {
	var notices []models.Notice
	if result := config.DB.Where("is_active = ?", true).Find(&notices); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(notices)
}

func CreateNotice(c *fiber.Ctx) error {
	notice := new(models.Notice)
	if err := c.BodyParser(notice); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if result := config.DB.Create(&notice); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(201).JSON(notice)
}

// Gallery
func GetGallery(c *fiber.Ctx) error {
	var gallery []models.Gallery
	if result := config.DB.Find(&gallery); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(gallery)
}

func AddToGallery(c *fiber.Ctx) error {
	item := new(models.Gallery)
	if err := c.BodyParser(item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if result := config.DB.Create(&item); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(201).JSON(item)
}

func UpdateGalleryItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item := new(models.Gallery)
	if err := c.BodyParser(item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if result := config.DB.Model(&models.Gallery{}).Where("id = ?", id).Updates(item); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(item)
}

func DeleteGalleryItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if result := config.DB.Delete(&models.Gallery{}, "id = ?", id); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.SendStatus(204)
}

func IncrementGalleryLikes(c *fiber.Ctx) error {
	id := c.Params("id")
	if result := config.DB.Model(&models.Gallery{}).Where("id = ?", id).Update("likes", gorm.Expr("likes + ?", 1)); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.SendStatus(200)
}

func IncrementGalleryDownloads(c *fiber.Ctx) error {
	id := c.Params("id")
	if result := config.DB.Model(&models.Gallery{}).Where("id = ?", id).Update("downloads", gorm.Expr("downloads + ?", 1)); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.SendStatus(200)
}
