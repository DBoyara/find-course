package router

import (
	"github.com/DBoyara/find-course/pkg/models"
	"github.com/DBoyara/find-course/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
)

// USER handles all the user routes
var CALC fiber.Router

const collectionName = "calculators"

// SetupCalcRoutes
func SetupCalcRoutes() {
	CALC.Post("/", CreateCalculator)
	CALC.Get("/", GetCalculators)
	CALC.Get("/:id", GetCalculator)
	CALC.Put("/:id", UpdatetCalculator)
	CALC.Delete("/:id", DeleteCalculator)
}

func CreateCalculator(c *fiber.Ctx) error {
	collection := repository.MG.Db.Collection(collectionName)

	calc := new(models.UserCalculation)
	if err := c.BodyParser(calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	insertionResult, err := collection.InsertOne(c.Context(), calc)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdCalc := &models.UserCalculation{}
	createdRecord.Decode(createdCalc)

	return c.Status(201).JSON(createdCalc)
}

func GetCalculators(c *fiber.Ctx) error {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := repository.MG.Db.Collection(collectionName).Find(c.Context(), query)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var calculators []models.UserCalculation = make([]models.UserCalculation, 0)
	if err := cursor.All(c.Context(), &calculators); err != nil {
		return c.Status(500).SendString(err.Error())

	}

	return c.JSON(calculators)
}

func GetCalculator(c *fiber.Ctx) error {
	var resultUserCalculation models.UserCalculation
	calcId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	filter := bson.M{"_id": calcId}
	err = repository.MG.Db.Collection(collectionName).FindOne(c.Context(), filter).Decode(&resultUserCalculation)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(resultUserCalculation)
}

func UpdatetCalculator(c *fiber.Ctx) error {
	idParam := c.Params("id")
	calcId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	calc := new(models.UserCalculation)
	if err := c.BodyParser(calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: calcId}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "client_name", Value: calc.ClientName},
			},
		},
	}
	err = repository.MG.Db.Collection(collectionName).FindOneAndUpdate(c.Context(), filter, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(calc)
}

func DeleteCalculator(c *fiber.Ctx) error {
	calcId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: calcId}}
	result, err := repository.MG.Db.Collection(collectionName).DeleteOne(c.Context(), &query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if result.DeletedCount < 1 {
		return c.Status(404).SendString(ErrCalcDoesNotExist.Error())
	}

	// the record was deleted
	return c.SendStatus(204)
}
