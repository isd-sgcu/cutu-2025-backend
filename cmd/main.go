package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/isd-sgcu/cutu2025-backend/config"
	_ "github.com/isd-sgcu/cutu2025-backend/docs"
	"github.com/isd-sgcu/cutu2025-backend/infrastructure"
	"github.com/isd-sgcu/cutu2025-backend/middleware"
	"github.com/isd-sgcu/cutu2025-backend/repository"
	"github.com/isd-sgcu/cutu2025-backend/routes"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()
	app.Use(middleware.RequestLoggerMiddleware())

	// Connect to the database
	db := infrastructure.ConnectDatabase(cfg)

	// Initialize repositories
	repo := repository.NewUserRepository(db)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(repo)

	// Register routes
	routes.RegisterUserRoutes(app, userUsecase) // Register the user routes

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json", // URL to access the Swagger docs
	}))

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
