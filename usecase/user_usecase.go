package usecase

import (
	"fmt"

	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/infrastructure"
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

func (u *UserUsecase) Register(user *domain.User, imagePath string) (domain.TokenResponse, error) {
	// Upload the image to S3 and get the image URL
	imageKey := fmt.Sprintf("user-images/%s.jpg", user.ID) // Customize the key as needed
	imageURL, err := infrastructure.UploadFileToS3(imagePath, imageKey)
	if err != nil {
		return domain.TokenResponse{}, fmt.Errorf("failed to upload image: %v", err)
	}

	// Update user with the uploaded image URL
	user.ImageURL = imageURL

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
