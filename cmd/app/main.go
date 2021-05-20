package main

import (
	"log"
	"os"
	"time"

	"github.com/DBoyara/find-course/pkg/repository"
	"github.com/DBoyara/find-course/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var configDefaultCORS = cors.Config{
	AllowOrigins: "*",
	AllowMethods: "GET,POST,PUT,DELETE",
	AllowHeaders: "*",
}

var configDefaultLogger = fiberLogger.Config{
	Format:       "${red}[${time}] ${green}${status} - ${blue}${latency} ${method} ${path}${reset}\n",
	TimeFormat:   "15:04:05",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
	Output:       os.Stderr,
}

// CreateServer creates a new Fiber instance
func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {
	// Connect to Postgres
	repository.ConnectToDB()
	repository.ConnectToMongoDB()

	app := CreateServer()

	app.Use(requestid.New())
	app.Use(cors.New(configDefaultCORS))
	app.Use(fiberLogger.New(configDefaultLogger))

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
