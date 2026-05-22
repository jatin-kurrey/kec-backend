package handlers

import (
	"kec-backend/internal/config"
	"kec-backend/internal/models"
	"kec-backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var admin models.Admin
	if result := config.DB.Where("username = ?", req.Username).First(&admin); result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if !utils.CheckPasswordHash(req.Password, admin.PasswordHash) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(admin.Username, admin.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token, "user": admin})
}

// For initial setup, allow creating a super admin if none exist
func CreateInitialAdmin(c *fiber.Ctx) error {
	var count int64
	config.DB.Model(&models.Admin{}).Count(&count)
	if count > 0 {
		return c.Status(403).JSON(fiber.Map{"error": "Admin already exists"})
	}

	var admin models.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// In a real app, this would be a password from the request
	password := "kec_admin_2026" // Default initial password
	hash, _ := utils.HashPassword(password)
	admin.PasswordHash = hash
	admin.Role = "SUPER_ADMIN"

	if result := config.DB.Create(&admin); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Initial admin created", "username": admin.Username, "password": password})
}

func GetAdmins(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "SUPER_ADMIN" {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied: Super Administrator privileges required"})
	}

	var admins []models.Admin
	if err := config.DB.Order("created_at desc").Find(&admins).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch admins"})
	}

	return c.JSON(admins)
}

type CreateAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func CreateAdmin(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "SUPER_ADMIN" {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied: Super Administrator privileges required"})
	}

	var req CreateAdminRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if req.Username == "" || req.Password == "" || req.Role == "" {
		return c.Status(400).JSON(fiber.Map{"error": "All fields are required"})
	}

	// Verify unique username
	var count int64
	config.DB.Model(&models.Admin{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Username already exists"})
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	admin := models.Admin{
		Username:     req.Username,
		PasswordHash: hash,
		Role:         req.Role,
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(admin)
}

func DeleteAdmin(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	username := c.Locals("username").(string)
	if role != "SUPER_ADMIN" {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied: Super Administrator privileges required"})
	}

	id := c.Params("id")
	var adminToDelete models.Admin
	if err := config.DB.Where("id = ?", id).First(&adminToDelete).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Admin account not found"})
	}

	// Prevent deleting oneself
	if adminToDelete.Username == username {
		return c.Status(400).JSON(fiber.Map{"error": "You cannot delete your own account"})
	}

	if err := config.DB.Delete(&adminToDelete).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete admin account"})
	}

	return c.JSON(fiber.Map{"message": "Admin account deleted successfully"})
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

func ChangeSelfPassword(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var req UpdatePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if req.NewPassword == "" {
		return c.Status(400).JSON(fiber.Map{"error": "New password is required"})
	}

	if len(req.NewPassword) < 6 {
		return c.Status(400).JSON(fiber.Map{"error": "Password must be at least 6 characters long"})
	}

	var admin models.Admin
	if err := config.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Admin user not found"})
	}

	if !utils.CheckPasswordHash(req.CurrentPassword, admin.PasswordHash) {
		return c.Status(400).JSON(fiber.Map{"error": "Incorrect current password"})
	}

	newHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash new password"})
	}

	admin.PasswordHash = newHash
	if err := config.DB.Save(&admin).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}

type AdminPasswordResetRequest struct {
	NewPassword string `json:"new_password"`
}

func ResetAdminPassword(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "SUPER_ADMIN" {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied: Super Administrator privileges required"})
	}

	id := c.Params("id")
	var req AdminPasswordResetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if req.NewPassword == "" {
		return c.Status(400).JSON(fiber.Map{"error": "New password is required"})
	}

	if len(req.NewPassword) < 6 {
		return c.Status(400).JSON(fiber.Map{"error": "Password must be at least 6 characters long"})
	}

	var admin models.Admin
	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Admin account not found"})
	}

	newHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash new password"})
	}

	admin.PasswordHash = newHash
	if err := config.DB.Save(&admin).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password updated successfully"})
}
