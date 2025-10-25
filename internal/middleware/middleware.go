package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Example: add request ID or logging middleware here
func RequestID(c *fiber.Ctx) error {
	c.Set("X-Request-ID", "req-"+uuid.New().String())
	return c.Next()
}
