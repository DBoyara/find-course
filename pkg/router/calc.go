package router

import (
	"log"

	"github.com/DBoyara/find-course/pkg/models"
	"github.com/DBoyara/find-course/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

// USER handles all the user routes
var CALC fiber.Router

// SetupCalcRoutes
func SetupCalcRoutes() {
	CALC.Post("/", CreateCalculator)
	CALC.Get("/", GetCalculators)
	CALC.Get("/:id", GetCalculator)
	CALC.Put("/:id", UpdatetCalculator)
	CALC.Delete("/:id", DeleteCalculator)
}

func CreateCalculator(c *fiber.Ctx) error {
	collection := repository.MG.Db.Collection("calculators")

	calc := new(models.UserCalculation)
	if err := c.BodyParser(calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	log.Println(collection.Name())

	insertionResult, err := collection.InsertOne(c.Context(), calc)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdCalc := &models.UserCalculation{}
	createdRecord.Decode(createdCalc)

	// return the created Employee in JSON format
	return c.Status(201).JSON(createdCalc)
}

func GetCalculators(c *fiber.Ctx) error {
	return nil
}

func GetCalculator(c *fiber.Ctx) error {
	log.Println(c, "sss%s", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": c.Params("id"),
	})
}

func UpdatetCalculator(c *fiber.Ctx) error {
	return nil
}

func DeleteCalculator(c *fiber.Ctx) error {
	return nil
}
