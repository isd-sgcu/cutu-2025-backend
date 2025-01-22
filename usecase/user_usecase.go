package usecase

import (
	"fmt"

	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/utils"
)

type UserUsecase struct {
	Repo UserRepositoryInterface
}

type UserRepositoryInterface interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetById(id string) (domain.User, error)
	Update(id string, user *domain.User) error
	Delete(id string) error
}

func NewUserUsecase(repo UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{Repo: repo}
}

// assing role based on phone number
func (u *UserUsecase) assignRole(user *domain.User) {
	// mock phone number
	staffPhones := []string{"06", "08", "09"}
	adminPhones := []string{"00", "07"}
	user.Role = domain.Member // Default role is Member

	// Check if the user's phone number matches the staff numbers
	if user.Phone != "" {
		for _, phone := range staffPhones {
			if user.Phone == phone {
				user.Role = domain.Staff
				break
			}
		}

		// Check if the user's phone number matches the admin numbers
		for _, phone := range adminPhones {
			if user.Phone == phone {
				user.Role = domain.Admin // Set role to Admin
				break
			}
		}
	}
}

func (u *UserUsecase) Register(user *domain.User) (domain.TokenResponse, error) {
	user.IsEntered = false
	u.assignRole(user)
	// Save user in the repository
	if err := u.Repo.Create(user); err != nil {
		return domain.TokenResponse{}, err
	}

	// Generate access and refresh tokens
	jwtSecret := utils.GetEnv("SECRET_JWT_KEY", "")
	accessToken, err := utils.GenerateTokens(user.ID, jwtSecret)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	return domain.TokenResponse{
		UserID:      user.ID,
		AccessToken: accessToken,
	}, nil
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.Repo.GetAll()
}

func (u *UserUsecase) GetById(id string) (domain.User, error) {
	return u.Repo.GetById(id)
}

func (u *UserUsecase) Update(id string, updatedUser *domain.User) error {
	_, err := u.Repo.GetById(id)
	if err != nil {
		return err
	}
	return u.Repo.Update(id, updatedUser)
}

// adjust IsEntered from False to True for scanning qrcode
func (u *UserUsecase) ScanQR(id string) (domain.User, error) {
	user, err := u.Repo.GetById(id)
	if err != nil {
		return domain.User{}, err
	}

	// Return early if the user has already entered
	if user.IsEntered {
		return domain.User{}, domain.ErrUserAlreadyEntered
	}

	// Update only the necessary field instead of creating a new object
	user.IsEntered = true

	err = u.Repo.Update(id, &user)

	return user, err
}

// adjust role
func (u *UserUsecase) UpdateRole(id string, role domain.Role) error {
	user, err := u.Repo.GetById(id)
	if err != nil {
		return err
	}
	user.Role = role
	return u.Repo.Update(id, &user)
}

// Get QR code URL
func (u *UserUsecase) GetQRURL(id string) (string, error) {
	user, err := u.Repo.GetById(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://localhost:4000/api/users/qr/%s", user.ID), nil
}

// Delete user
func (u *UserUsecase) Delete(id string) error {
	return u.Repo.Delete(id)
}
