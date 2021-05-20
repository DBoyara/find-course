package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDB connects the server with database
func ConnectToMongoDB() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}

	dsn := fmt.Sprintf("mongodb://%s:%s@localhost:%s/%s",
		os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASS"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"),
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to MongoDB...")
	database := client.Database(os.Getenv("MONGO_DB"))
	log.Printf("Connected to MongoDB %s", database.Name())
}
