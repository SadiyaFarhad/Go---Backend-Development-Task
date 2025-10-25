package routes

import (
	"go-backend-user/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// ✅ Default root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber is running!")
	})

	// Your existing routes
	app.Post("/users", handler.CreateUser)
	app.Get("/users/:id", handler.GetUser)
	app.Get("/users", handler.ListUsers)
	app.Put("/users/:id", handler.UpdateUser)
	app.Delete("/users/:id", handler.DeleteUser)
}
