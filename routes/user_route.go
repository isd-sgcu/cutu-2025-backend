package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/handler"
	"github.com/isd-sgcu/cutu2025-backend/infrastructure"
	"github.com/isd-sgcu/cutu2025-backend/middleware"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
)

func RegisterUserRoutes(app *fiber.App, userUsecase *usecase.UserUsecase, s3 *infrastructure.S3Client) {
	userHandler := handler.NewUserHandler(userUsecase, s3)

	api := app.Group("/api/users")

	api.Get("/",
		middleware.RoleMiddleware(
			string(domain.Staff),
			string(domain.Admin),
		),
		userHandler.GetAll)

	api.Get("/:id", middleware.AuthMiddleware(), userHandler.GetById)

	api.Post("/register", userHandler.Register)

	api.Patch("/:id", middleware.RoleMiddleware(
		string(domain.Admin),
	),
		userHandler.Update)
}
