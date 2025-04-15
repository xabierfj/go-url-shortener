// cmd/server/main.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/xabierfj/go-url-shortener/internal/database"
	"github.com/xabierfj/go-url-shortener/internal/handlers"
	"log"
)

func main() {
	// Load .env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Start Fiber app
	app := fiber.New()
	handlers.SetupRoutes(app, db)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
