package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Remove trailing slash
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization", // Add Authorization here
		AllowCredentials: true,
	}))

	// Connect to the database
	db := infrastructure.ConnectDatabase(cfg)

	// Connect to S3
	s3 := infrastructure.ConnectToS3(cfg)

	// Connect to Cache
	cache := infrastructure.ConnectRedis(cfg)

	// Initialize repositories
	repo := repository.NewUserRepository(db)
	storage := repository.NewStorageRepository(s3)
	cacheRepo := repository.NewCacheRepository(cache)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(repo, storage, cacheRepo)

	// Register routes
	routes.RegisterUserRoutes(app, userUsecase) // Register the user routes

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json", // URL to access the Swagger docs
	}))

	// Start the server
	if err := app.Listen(":4000"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
