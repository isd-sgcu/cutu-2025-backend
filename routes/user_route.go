package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/handler"
	"github.com/isd-sgcu/cutu2025-backend/middleware"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
)

func RegisterUserRoutes(app *fiber.App, userUsecase *usecase.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)

	api := app.Group("/api/users")

	api.Get("/",
		middleware.RoleMiddleware(
			string(domain.RoleAdmin),
			string(domain.RoleStaff),
		),
		userHandler.GetAll)

	api.Get("/:id", middleware.AuthMiddleware(), userHandler.GetById)

	api.Post("/register", userHandler.Register)

	api.Patch("/:id", middleware.RoleMiddleware(
		string(domain.RoleAdmin),
	),
		userHandler.Update)
}
