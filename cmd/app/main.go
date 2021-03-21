package main

import (
	"log"
	"net"
	"os"

	"github.com/DBoyara/find-course/pkg/repository"
	"github.com/DBoyara/find-course/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CreateServer creates a new Fiber instance
func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

const (
	defaultHost = "127.0.0.1"
	defaultPort = "3000"
	defaultUser = "db_user"
	defaultPass = "pass"
	defaultDB = "db"
	defaultDBHost = "0.0.0.0"
)

func main() {
	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	dbHost, ok := os.LookupEnv("PSQL_HOST")
	if !ok {
		dbHost = defaultDBHost
	}

	dbUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		dbUser = defaultUser
	}

	dbPass, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		dbPass = defaultPass
	}

	db, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		db = defaultDB
	}

	if err := execute(net.JoinHostPort(host, port), dbHost, dbUser, dbPass, db); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(appAddress string, host string, dbUser string, dbPass string, db string) (err error) {
	// Connect to Postgres
	err = repository.ConnectToDB(host, dbUser, dbPass, db)
	if err != nil {
		log.Fatal("Fail with database. \n", err)
		return err
	}

	app := CreateServer()

	app.Use(cors.New())

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app.Listen(appAddress)
}
