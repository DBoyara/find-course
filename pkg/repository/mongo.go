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

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MG MongoInstance

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

	log.Println("Connecting to MongoDB...")
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(os.Getenv("MONGO_DB"))
	log.Printf("Connected to MongoDB %s", database.Name())

	MG = MongoInstance{
		Client: client,
		Db:     database,
	}

}
