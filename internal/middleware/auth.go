package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	// Skip auth for OPTIONS requests (CORS preflight)
	if c.Method() == "OPTIONS" {
		return c.Next()
	}

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing authorization header"})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	username, _ := claims["username"].(string)
	role, _ := claims["role"].(string)

	c.Locals("username", username)
	c.Locals("role", role)

	// Role-Based Access Control:
	// GALLERY_ADMIN is ONLY allowed to access gallery-related endpoints or file uploads.
	if role == "GALLERY_ADMIN" {
		path := c.Path()
		isGalleryRoute := strings.Contains(path, "/admin/gallery") || strings.Contains(path, "/admin/upload")
		if !isGalleryRoute {
			return c.Status(403).JSON(fiber.Map{"error": "Access denied: Gallery administrator permissions only"})
		}
	}

	return c.Next()
}
