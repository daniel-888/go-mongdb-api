package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func defaltHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("all good")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	MONGO_USERNAME := os.Getenv("MONGO_USERNAME")
	MONGO_PASSWORD := os.Getenv("MONGO_PASSWORD")
	MONGO_HOSTNAME := os.Getenv("MONGO_HOSTNAME")
	MONGO_PORT := os.Getenv("MONGO_PORT")
	MONGO_DB := os.Getenv("MONGO_DB")

	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + MONGO_USERNAME + ":" + MONGO_PASSWORD + "@" + MONGO_HOSTNAME + ":" + MONGO_PORT + "/" + MONGO_DB + "?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("mongodb connected")
	defer client.Disconnect(ctx)

	// collection := client.Database(MONGO_DB).Collection("users")

	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())

	app.Get("/", defaltHandler)

	// get the port
	port := os.Getenv("PORT")
	
	// launch the app
	launchError := app.Listen(":" + port)
	if launchError != nil {
		panic(launchError)
	}
}
