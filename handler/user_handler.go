package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: usecase}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := h.Usecase.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.Usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
