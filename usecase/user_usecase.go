package usecase

import (
	"bytes"
	"fmt"
	"time"

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
	GetByPhone(phone string) (domain.User, error)
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

// Check Scanning is same day
func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (u *UserUsecase) Register(user *domain.User, storage *infrastructure.S3Client, fileBytes []byte, fileName string) (domain.TokenResponse, error) {
	u.assignRole(user)

	// Upload the file using the existing UploadFile method
	fileReader := bytes.NewReader(fileBytes)
	s3Key := fmt.Sprintf("cutu-2025/%s", fileName)
	s3URL, err := storage.UploadFile(utils.GetEnv("S3_BUCKET_NAME", ""), s3Key, fileReader)

	if err != nil {
		return domain.TokenResponse{}, err
	}

	user.ImageURL = s3URL
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

func (u *UserUsecase) SignIn(id string) (domain.TokenResponse, error) {
	user, err := u.Repo.GetById(id)
	if err != nil {
		return domain.TokenResponse{}, err
	}

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

	now := time.Now()
	// Check if the user has already entered today
	if user.LastEntered != nil && isSameDay(*user.LastEntered, now) {
		return user, domain.ErrUserAlreadyEntered
	}

	// Update the last entry timestamp
	user.LastEntered = &now
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

// Get Card ID image
func (u *UserUsecase) GetCardID(id string) (string, error) {
	user, err := u.Repo.GetById(id)
	if err != nil {
		return "", err
	}
	return user.ImageURL, nil
}

// Add Staff by phone number
func (u *UserUsecase) AddStaff(phone string) error {
	user, err := u.Repo.GetByPhone(phone)
	if err != nil {
		return err
	}
	if user.Role == domain.Staff {
		return domain.ErrUserAlreadyStaff
	}
	user.Role = domain.Staff
	return u.Repo.Update(user.ID, &user)
}
