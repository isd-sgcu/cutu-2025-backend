package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/config"
	"github.com/isd-sgcu/cutu2025-backend/infrastructure"
	"github.com/isd-sgcu/cutu2025-backend/repository"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
	"github.com/isd-sgcu/cutu2025-backend/routes" // Assuming you have routes package
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Connect to the database
	db := infrastructure.ConnectDatabase(cfg)

	// Initialize repositories
	repo := repository.NewUserRepository(db)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(repo)

	// Register routes
	routes.RegisterUserRoutes(app, userUsecase) // Register the user routes

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
