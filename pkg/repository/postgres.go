package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/DBoyara/find-course/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB represents a Database instance
var DB *gorm.DB

// ConnectToDB connects the server with database
func ConnectToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}
	log.Println(os.Getenv("DB_HOST"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"),
	)

	log.Print("Connecting to PostgreSQL DB...")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")

	log.Print("Running the migrations...")
	err = DB.AutoMigrate(&models.User{}, &models.Claims{})
	if err != nil {
		log.Fatal("Failed to auto-migrate. \n", err)
		os.Exit(2)
	}

}
