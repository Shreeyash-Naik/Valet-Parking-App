package api

import (
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Get("/checkin", Checkin)
		api.Get("/checkout", Checkout)
		api.Get("/search", Search)
	}
}
