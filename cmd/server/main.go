package main

import (
	"go-backend-user/internal/logger"
	"go-backend-user/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber is running!")
	})

	routes.SetupRoutes(app)

	logger.Log.Info("Server starting on port 3000")
	app.Listen(":3000")
}
