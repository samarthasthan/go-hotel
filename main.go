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

// these are dummy credentials use your own production credentails and store them in .env file
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
	apiv1 := app.Group("/api/v1/")

	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Get("/users", userHandler.HandleGetUsers)

	// Listen fiber routes
	app.Listen(":8000")
}
