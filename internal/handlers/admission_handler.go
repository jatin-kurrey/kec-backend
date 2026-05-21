package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// --- Public Handlers ---

func GetAdmissionGuide(c *fiber.Ctx) error {
	var guide models.AdmissionGuide
	if err := config.DB.First(&guide).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Guide not found"})
	}
	return c.JSON(guide)
}

func GetAdmissionSteps(c *fiber.Ctx) error {
	var steps []models.AdmissionStep
	config.DB.Order("sort_order").Find(&steps)
	return c.JSON(steps)
}

func GetAdmissionEligibility(c *fiber.Ctx) error {
	var items []models.AdmissionEligibility
	config.DB.Order("sort_order").Find(&items)
	return c.JSON(items)
}

func GetAdmissionDocuments(c *fiber.Ctx) error {
	var docs []models.AdmissionDocument
	config.DB.Order("sort_order").Find(&docs)
	return c.JSON(docs)
}

func GetAdmissionFees(c *fiber.Ctx) error {
	var fees []models.AdmissionFee
	config.DB.Order("sort_order").Find(&fees)
	return c.JSON(fees)
}

// --- Admin Handlers ---

// Guide
func UpdateAdmissionGuide(c *fiber.Ctx) error {
	var guide models.AdmissionGuide
	if err := config.DB.First(&guide).Error; err != nil {
		// Create if not exists
		if err := c.BodyParser(&guide); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		config.DB.Create(&guide)
		return c.JSON(guide)
	}
	if err := c.BodyParser(&guide); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&guide)
	return c.JSON(guide)
}

// Steps
func CreateAdmissionStep(c *fiber.Ctx) error {
	step := new(models.AdmissionStep)
	if err := c.BodyParser(step); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&step)
	return c.JSON(step)
}

func UpdateAdmissionStep(c *fiber.Ctx) error {
	id := c.Params("id")
	var step models.AdmissionStep
	if err := config.DB.First(&step, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Step not found"})
	}
	if err := c.BodyParser(&step); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&step)
	return c.JSON(step)
}

func DeleteAdmissionStep(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.AdmissionStep{}, id)
	return c.SendStatus(204)
}

// Eligibility
func CreateAdmissionEligibility(c *fiber.Ctx) error {
	item := new(models.AdmissionEligibility)
	if err := c.BodyParser(item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&item)
	return c.JSON(item)
}

func UpdateAdmissionEligibility(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.AdmissionEligibility
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&item)
	return c.JSON(item)
}

func DeleteAdmissionEligibility(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.AdmissionEligibility{}, id)
	return c.SendStatus(204)
}

// Documents
func CreateAdmissionDocument(c *fiber.Ctx) error {
	doc := new(models.AdmissionDocument)
	if err := c.BodyParser(doc); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&doc)
	return c.JSON(doc)
}

func UpdateAdmissionDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	var doc models.AdmissionDocument
	if err := config.DB.First(&doc, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Document not found"})
	}
	if err := c.BodyParser(&doc); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&doc)
	return c.JSON(doc)
}

func DeleteAdmissionDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.AdmissionDocument{}, id)
	return c.SendStatus(204)
}

// Fees
func CreateAdmissionFee(c *fiber.Ctx) error {
	fee := new(models.AdmissionFee)
	if err := c.BodyParser(fee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&fee)
	return c.JSON(fee)
}

func UpdateAdmissionFee(c *fiber.Ctx) error {
	id := c.Params("id")
	var fee models.AdmissionFee
	if err := config.DB.First(&fee, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Fee not found"})
	}
	if err := c.BodyParser(&fee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&fee)
	return c.JSON(fee)
}

func DeleteAdmissionFee(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.AdmissionFee{}, id)
	return c.SendStatus(204)
}
