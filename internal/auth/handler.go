package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
	}
	if err := c.BodyParser(&req); err != nil || req.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid username"})
	}

	token, err := GenerateToken(req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "token generation failed"})
	}
	return c.JSON(fiber.Map{"token": token})
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "no token provided"})
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	c.Locals("username", claims.Username)
	return c.Next()
}

func ProfileHandler(c *fiber.Ctx) error {
	username := c.Locals("username")
	return c.JSON(fiber.Map{
		"user":      username,
		"message":   "Welcome to your profile!",
		"timestamp": time.Now(),
	})
}
