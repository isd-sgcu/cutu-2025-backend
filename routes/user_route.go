package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/handler"
	"github.com/isd-sgcu/cutu2025-backend/middleware"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
)

func RegisterUserRoutes(app *fiber.App, userUsecase *usecase.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)

	api := app.Group("/api/users", middleware.TimingMiddleware())
	api.Get("/", userHandler.GetAll)
	api.Post("/register", middleware.AuthMiddleware(), userHandler.Register)
}
