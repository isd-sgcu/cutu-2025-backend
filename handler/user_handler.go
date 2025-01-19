package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
)

// UserHandler represents the handler for user-related endpoints
type UserHandler struct {
	Usecase *usecase.UserUsecase
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: usecase}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param user body domain.User true "User data"
// @Success 201 {object} domain.User
// @Failure 400 {object} domain.ErrorResponse "Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 500 {object} domain.ErrorResponse "Failed to create user"
// @Router /api/users/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}
	if err := h.Usecase.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetAll godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Produce  json
// @Success 200 {array} domain.User
// @Failure 500 {object} domain.ErrorResponse "Failed to fetch users"
// @Router /api/users [get]
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.Usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to fetch users"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
