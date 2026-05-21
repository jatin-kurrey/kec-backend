package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}

	// Generate unique filename
	uniqueID := uuid.New().String()
	filename := fmt.Sprintf("%d-%s-%s", time.Now().Unix(), uniqueID, file.Filename)

	// Save file to uploads directory
	if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Return file URL
	// Note: In production, you'd use your domain
	fileURL := fmt.Sprintf("/uploads/%s", filename)
	return c.JSON(fiber.Map{"url": fileURL})
}
