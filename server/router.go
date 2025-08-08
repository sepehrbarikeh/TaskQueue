package server




import (

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h Handler) {
	app.Post("/enqueue", h.EnqueueJob)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
