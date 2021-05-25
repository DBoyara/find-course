package router

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrCalcDoesNotExist = errors.New("calc does not exist")
)

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	USER = api.Group("/user")
	SetupUserRoutes()

	CALC = api.Group("/calculators")
	SetupCalcRoutes()
}
