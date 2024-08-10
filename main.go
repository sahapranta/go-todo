package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	client := connection()

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB. ", err)
	}

	collection = client.Database("go_todo").Collection("todos")

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", deleteTodos)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	log.Fatal(app.Listen(":" + PORT))
}

func connection() *mongo.Client {
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB. ", err)
	}

	return client
}
