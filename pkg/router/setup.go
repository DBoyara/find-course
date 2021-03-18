package router

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	USER = api.Group("/user")
	SetupUserRoutes()
}
