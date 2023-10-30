package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/samarthasthan/go-hotel/api"
	"github.com/samarthasthan/go-hotel/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const DBUri = "mongodb+srv://xinvo:xinvopassword@xinvodb.fzm3svb.mongodb.net/?retryWrites=true&w=majority"
const DBName = "xinvodb"
const DBColl = "users"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	// MongoDB connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBUri))
	if err != nil {
		log.Fatal(err)
	}

	// Handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	// New fiber instance
	app := fiber.New(config)

	app.Get("/user/:id", userHandler.HandleGetUser)
	app.Get("/users", userHandler.HandleGetUsers)

	// Listen fiber routes
	app.Listen(":8000")
}
