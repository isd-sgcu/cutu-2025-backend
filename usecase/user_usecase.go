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
}

func NewUserUsecase(repo UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{Repo: repo}
}

// assing role based on phone number
func (u *UserUsecase) assignRole(user *domain.User) {
	// mock phone number
	staff_phones := []string{"06", "08", "09"}
	user.Role = domain.Student
	if user.Phone != "" {
		for _, phone := range staff_phones {
			if user.Phone == phone {
				user.Role = domain.Staff
				break
			}
		}
	}
}

func (u *UserUsecase) Register(user *domain.User) (domain.TokenResponse, error) {
	u.assignRole(user)
	// Save user in the repository
	if err := u.Repo.Create(user); err != nil {
		return domain.TokenResponse{}, err
	}

	// Generate access and refresh tokens
	jwtSecret := utils.GetEnv("SECRET_JWT_KEY", "")
	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, string(user.Role), jwtSecret)
	if err != nil {
		return domain.TokenResponse{}, err
	}

	return domain.TokenResponse{
		UserID:       user.ID,
		QrURL:        fmt.Sprintf("http://localhost:4000/api/users/qr/%s", user.ID),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
