package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TimingMiddleware logs the request duration
func TimingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next() // Continue to the next middleware or handler
		duration := time.Since(start)
		log.Printf("Request took %v", duration)
		return err
	}
}

// AuthMiddleware checks for a dummy authorization header
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader != "Bearer my-secret-token" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		return c.Next() // Continue if authorized
	}
}
