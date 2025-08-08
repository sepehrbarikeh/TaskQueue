package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h Handler) {
	// API routes
	api := app.Group("/api/v1")
	api.Post("/jobs", h.EnqueueJob)
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"time":   time.Now().UTC(),
		})
	})

	// Root health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
