package handler

import (
	"errors"
	"strings"

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
// @Param id formData string true "User ID"
// @Param name formData string true "User Name"
// @Param email formData string true "User Email"
// @Param phone formData string true "User Phone"
// @Param university formData string true "User University"
// @Param sizeJersey formData string true "Jersey Size"
// @Param foodLimitation formData string false "Food Limitation"
// @Param invitationCode formData string false "Invitation Code"
// @Param status formData domain.Status true "User Status"
// @Param image formData file true "User Image"
// @Param graduatedYear formData string false "Graduated Year"
// @Param faculty formData string false "Faculty"
// @Param education formData domain.Education true "Education"
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
	// ! Cant use S3 for now
	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to read image file"})
	// }

	// Upload the file using the existing UploadFile method
	// ! Cant use S3 for now
	// fileReader := bytes.NewReader(fileBytes)
	// s3Key := fmt.Sprintf("cutu-2025/%s", imageFile.Filename)
	// s3URL, err := h.S3Service.UploadFile(utils.GetEnv("S3_BUCKET_NAME", ""), s3Key, fileReader)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to upload image to S3 " + err.Error()})
	// }

	// Map form values to user object
	user := &domain.User{
		ID:             form.Value["id"][0],
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
		Status: domain.StatusAlumni,
		GraduatedYear: func() *string {
			if v := form.Value["graduatedYear"]; len(v) > 0 {
				return &v[0]
			}
			return nil
		}(),
		Faculty: func() *string {
			if v := form.Value["faculty"]; len(v) > 0 {
				return &v[0]
			}
			return nil
		}(),
		Education: domain.Education(form.Value["education"][0]),
		ImageURL:  "https://picsum.photos/id/237/200/300", //TODO: Waiting for S3
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
// @Failure 403 {object} domain.ErrorResponse "Forbidden"
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

// GetById godoc
// @Summary Scan QR code
// @Description Retrieve a user by its ID
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 400 {object} domain.ErrorResponse "User has already entered"
// @Router /api/users/qr/{id} [post]
func (h *UserHandler) ScanQR(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.Usecase.ScanQR(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyEntered) {
			return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "User has already entered"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// Change Role godoc
// @Summary Update user role by ID
// @Description Update a user by its ID
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param role body domain.Role true "User Role"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse "Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 403 {object} domain.ErrorResponse "Forbidden"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to update user role"
// @Router /api/users/role/{id} [patch]
func (h *UserHandler) UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")
	role := new(domain.Role)
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}
	if err := h.Usecase.UpdateRole(id, *role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to update this role user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Change Role godoc
// @Summary Update user role by ID
// @Description Update a user by its ID
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param user body domain.User true "User data"
// @Success 204
// @Failure 400 {object} domain.ErrorResponse "Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 403 {object} domain.ErrorResponse "Forbidden"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to update user role"
// @Router /api/users [patch]
func (h *UserHandler) UpdateMyAccountInfo(c *fiber.Ctx) error {
	role := new(domain.Role)
	token := strings.Split(c.Get("Authorization"), "Bearer")
	id, err := utils.DecodeToken(token[1], utils.GetEnv("SECRET_JWT_KEY", ""))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.ErrorResponse{Error: "Unauthorized"})
	}

	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}
	if err := h.Usecase.UpdateRole(id, *role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to update this role user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetQRURL godoc
// @Summary Get QR code URL
// @Description Retrieve a QR code URL for a user
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} domain.QrResponse
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to fetch user"
// @Router /api/users/qr/{id} [get]
func (h *UserHandler) GetQRURL(c *fiber.Ctx) error {
	id := c.Params("id")
	qrURL, err := h.Usecase.GetQRURL(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(domain.ErrorResponse{Error: "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(domain.QrResponse{QrURL: qrURL})
}

// Delete godoc
// @Summary Delete user by ID
// @Description Delete a user by its ID
// @Produce  json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 204
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 403 {object} domain.ErrorResponse "Forbidden"
// @Failure 404 {object} domain.ErrorResponse "User not found"
// @Failure 500 {object} domain.ErrorResponse "Failed to delete user"
// @Router /api/users/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Usecase.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to delete user"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Signin godoc
// @Summary Signin
// @Description Signin
// @Produce  json
// @Param id body string true "User ID"
// @Success 200 {object} domain.TokenResponse
// @Failure 400 {object} domain.ErrorResponse "Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized"
// @Failure 500 {object} domain.ErrorResponse "Failed to signin"
// @Router /api/users/signin [post]
func (h *UserHandler) Signin(c *fiber.Ctx) error {
	id := new(string)
	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorResponse{Error: "Invalid input"})
	}

	tokenResponse, err := h.Usecase.SignIn(*id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrorResponse{Error: "Failed to signin"})
	}

	return c.Status(fiber.StatusOK).JSON(tokenResponse)
}
