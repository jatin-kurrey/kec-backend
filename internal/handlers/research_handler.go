package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

// --- Research Areas ---

func GetResearchAreas(c *fiber.Ctx) error {
	var areas []models.ResearchArea
	if err := config.DB.Where("is_active = ?", true).Order("created_at asc").Find(&areas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research areas"})
	}
	return c.JSON(areas)
}

func GetAllResearchAreasAdmin(c *fiber.Ctx) error {
	var areas []models.ResearchArea
	if err := config.DB.Order("created_at desc").Find(&areas).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research areas"})
	}
	return c.JSON(areas)
}

func CreateResearchArea(c *fiber.Ctx) error {
	area := new(models.ResearchArea)
	if err := c.BodyParser(area); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Create(area).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create research area"})
	}
	return c.Status(201).JSON(area)
}

func UpdateResearchArea(c *fiber.Ctx) error {
	id := c.Params("id")
	area := new(models.ResearchArea)
	if err := c.BodyParser(area); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Model(&models.ResearchArea{}).Where("id = ?", id).Updates(area).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update research area"})
	}
	return c.JSON(fiber.Map{"message": "Research area updated successfully"})
}

func DeleteResearchArea(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.ResearchArea{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete research area"})
	}
	return c.JSON(fiber.Map{"message": "Research area deleted successfully"})
}

// --- Research Projects ---

func GetResearchProjects(c *fiber.Ctx) error {
	var projects []models.ResearchProject
	if err := config.DB.Where("is_active = ?", true).Order("created_at desc").Find(&projects).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research projects"})
	}
	return c.JSON(projects)
}

func GetAllResearchProjectsAdmin(c *fiber.Ctx) error {
	var projects []models.ResearchProject
	if err := config.DB.Order("created_at desc").Find(&projects).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research projects"})
	}
	return c.JSON(projects)
}

func CreateResearchProject(c *fiber.Ctx) error {
	project := new(models.ResearchProject)
	if err := c.BodyParser(project); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Create(project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create research project"})
	}
	return c.Status(201).JSON(project)
}

func UpdateResearchProject(c *fiber.Ctx) error {
	id := c.Params("id")
	project := new(models.ResearchProject)
	if err := c.BodyParser(project); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Model(&models.ResearchProject{}).Where("id = ?", id).Updates(project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update research project"})
	}
	return c.JSON(fiber.Map{"message": "Research project updated successfully"})
}

func DeleteResearchProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.ResearchProject{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete research project"})
	}
	return c.JSON(fiber.Map{"message": "Research project deleted successfully"})
}

// --- Research Facilities ---

func GetResearchFacilities(c *fiber.Ctx) error {
	var facilities []models.ResearchFacility
	if err := config.DB.Where("is_active = ?", true).Order("created_at asc").Find(&facilities).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research facilities"})
	}
	return c.JSON(facilities)
}

func GetAllResearchFacilitiesAdmin(c *fiber.Ctx) error {
	var facilities []models.ResearchFacility
	if err := config.DB.Order("created_at desc").Find(&facilities).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research facilities"})
	}
	return c.JSON(facilities)
}

func CreateResearchFacility(c *fiber.Ctx) error {
	facility := new(models.ResearchFacility)
	if err := c.BodyParser(facility); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Create(facility).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create research facility"})
	}
	return c.Status(201).JSON(facility)
}

func UpdateResearchFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	facility := new(models.ResearchFacility)
	if err := c.BodyParser(facility); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Model(&models.ResearchFacility{}).Where("id = ?", id).Updates(facility).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update research facility"})
	}
	return c.JSON(fiber.Map{"message": "Research facility updated successfully"})
}

func DeleteResearchFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.ResearchFacility{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete research facility"})
	}
	return c.JSON(fiber.Map{"message": "Research facility deleted successfully"})
}

// --- Research Stats ---

func GetResearchStats(c *fiber.Ctx) error {
	var stats []models.ResearchStat
	if err := config.DB.Order("sort_order asc").Find(&stats).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch research stats"})
	}
	return c.JSON(stats)
}

func CreateResearchStat(c *fiber.Ctx) error {
	stat := new(models.ResearchStat)
	if err := c.BodyParser(stat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Create(stat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create research stat"})
	}
	return c.Status(201).JSON(stat)
}

func UpdateResearchStat(c *fiber.Ctx) error {
	id := c.Params("id")
	stat := new(models.ResearchStat)
	if err := c.BodyParser(stat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := config.DB.Model(&models.ResearchStat{}).Where("id = ?", id).Updates(stat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update research stat"})
	}
	return c.JSON(fiber.Map{"message": "Research stat updated successfully"})
}

func DeleteResearchStat(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.ResearchStat{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete research stat"})
	}
	return c.JSON(fiber.Map{"message": "Research stat deleted successfully"})
}
