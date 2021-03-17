package main

import (
	"log"

	"github.com/DBoyara/find-course/repository"
	"github.com/DBoyara/find-course/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CreateServer creates a new Fiber instance
func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	// Connect to Postgres
	repository.ConnectToDB()

	app := CreateServer()

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
