package repository

import (
	"fmt"
	"log"

	"github.com/DBoyara/find-course/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB represents a Database instance
var DB *gorm.DB

// ConnectToDB connects the server with database
func ConnectToDB(host string, user string, pass string, db string) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata", host, user, pass, db,
	)

	log.Print("Connecting to PostgreSQL DB...")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		return err
	}
	log.Println("connected")

	log.Print("Running the migrations...")
	err = DB.AutoMigrate(&models.User{}, &models.Claims{})
	if err != nil {
		log.Fatal("Failed to auto-migrate. \n", err)
		return err
	}

	return nil
}
