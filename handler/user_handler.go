package handler

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/infrastructure"
	"github.com/isd-sgcu/cutu2025-backend/usecase"
	"github.com/isd-sgcu/cutu2025-backend/utils"
)

// UserHandler represents the handler for user-related endpoints
type UserHandler struct {
	Usecase   *usecase.UserUsecase
	S3Service *infrastructure.S3Client
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(usecase *usecase.UserUsecase, s3Service *infrastructure.S3Client) *UserHandler {
	return &UserHandler{Usecase: usecase, S3Service: s3Service}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Accept  multipart/form-data
// @Produce  json
// @Security BearerAuth
// @Param name formData string true "User Name"
// @Param email formData string true "User Email"
// @Param phone formData string true "User Phone"
// @Param university formData string true "User University"
// @Param sizeJersey formData string true "Jersey Size"
// @Param foodLimitation formData string false "Food Limitation"
// @Param invitationCode formData string false "Invitation Code"
// @Param state formData string true "User State"
// @Param image formData file true "User Image"
// @Success 201 {object} domain.TokenResponse
// @Failure 400 {object} domain.ErrorResponse "Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 500 {object} domain.ErrorResponse "Failed to create user"
// @Router /api/users/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}

	// Get image file from form
	imageFiles := form.File["image"]
	if len(imageFiles) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Image is required"})
	}

	imageFile := imageFiles[0]

	// Open file stream
	file, err := imageFile.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to open image file"})
	}
	defer file.Close()

	// Convert the file stream to a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to read image file"})
	}

	// Convert byte slice to bytes.Reader for UploadFile method
	fileReader := bytes.NewReader(fileBytes)

	// Upload the file using the existing UploadFile method
	s3Key := fmt.Sprintf("cutu-2025/%s", imageFile.Filename)
	s3URL, err := h.S3Service.UploadFile(utils.GetEnv("S3_BUCKET_NAME", ""), s3Key, fileReader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to upload image to S3 " + err.Error()})
	}

	// Map form values to user object
	user := &domain.User{
		Name:           form.Value["name"][0],
		Email:          form.Value["email"][0],
		Phone:          form.Value["phone"][0],
		University:     form.Value["university"][0],
		SizeJersey:     form.Value["sizeJersey"][0],
		FoodLimitation: form.Value["foodLimitation"][0],
		InvitationCode: func() *string {
			if v := form.Value["invitationCode"]; len(v) > 0 {
				return &v[0]
			}
			return nil
		}(),
		State:    form.Value["state"][0],
		ImageURL: s3URL,
	}

	// Register user
	tokenResponse, err := h.Usecase.Register(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(tokenResponse)
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

// GetById godoc
// @Summary Get user by ID
// @Description Retrieve a user by its ID
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to fetch user"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.Usecase.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(domain.ErrorResponse{Error: "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// Update godoc
// @Summary Update user by ID
// @Description Update a user by its ID
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param user body domain.User true "User data"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse "Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to update user"
// @Router /api/users/{id} [patch]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}
	if err := h.Usecase.Update(id, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to update user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
