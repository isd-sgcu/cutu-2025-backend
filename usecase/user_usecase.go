package usecase

import (
	"bytes"
	"fmt"
	"time"

	"github.com/isd-sgcu/cutu2025-backend/domain"
	"github.com/isd-sgcu/cutu2025-backend/utils"
)

type UserUsecase struct {
	Repo    UserRepositoryInterface
	Storage StorageRepositoryInterface
}

type UserRepositoryInterface interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetById(id string) (domain.User, error)
	GetByPhone(phone string) (domain.User, error)
	GetByName(name string) ([]domain.User, error)
	IsUIDExists(uid string) (bool, error)
	Update(id string, user *domain.User) error
	Delete(id string) error
}

type StorageRepositoryInterface interface {
	UploadFile(bucketName, objectKey string, buffer *bytes.Reader) (string, error)
	DownloadFile(bucketName, objectKey, filePath string) error
	DeleteFile(bucketName, objectKey string) error
}

func NewUserUsecase(repo UserRepositoryInterface, storage StorageRepositoryInterface) *UserUsecase {
	return &UserUsecase{Repo: repo, Storage: storage}
}

func (u *UserUsecase) assignRole(user *domain.User) {
	staffPhones := []string{"06", "08", "09"}
	adminPhones := []string{"00", "07"}
	user.Role = domain.Member

	if user.Phone != "" {
		for _, phone := range staffPhones {
			if user.Phone == phone {
				user.Role = domain.Staff
				break
			}
		}

		for _, phone := range adminPhones {
			if user.Phone == phone {
				user.Role = domain.Admin
				break
			}
		}
	}
}

func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (u *UserUsecase) Register(user *domain.User, fileBytes []byte, fileName string) (domain.TokenResponse, error) {
	u.assignRole(user)

	for {
		user.UID = utils.GenerateUID()
		uidExists, err := u.Repo.IsUIDExists(user.UID)
		if err != nil {
			return domain.TokenResponse{}, fmt.Errorf("error checking UID uniqueness: %w", err)
		}
		if !uidExists {
			break
		}
	}

	fileReader := bytes.NewReader(fileBytes)
	s3Key := fmt.Sprintf("cutu-2025/%s", fileName)
	s3URL, err := u.Storage.UploadFile(utils.GetEnv("S3_BUCKET_NAME", ""), s3Key, fileReader)
	if err != nil {
		return domain.TokenResponse{}, fmt.Errorf("error uploading file: %w", err)
	}

	user.ImageURL = s3URL
	user.RegisteredAt = time.Now()

	if err := u.Repo.Create(user); err != nil {
		return domain.TokenResponse{}, fmt.Errorf("error saving user: %w", err)
	}

	jwtSecret := utils.GetEnv("SECRET_JWT_KEY", "")
	accessToken, err := utils.GenerateTokens(user.ID, jwtSecret)
	if err != nil {
		return domain.TokenResponse{}, fmt.Errorf("error generating tokens: %w", err)
	}

	return domain.TokenResponse{
		UserID:      user.ID,
		AccessToken: accessToken,
	}, nil
}

func (u *UserUsecase) GetAll(filter string) ([]domain.User, error) {
	if filter != "" {
		users, err := u.Repo.GetByName(filter)
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	return u.Repo.GetAll()
}

func (u *UserUsecase) GetById(id string) (domain.User, error) {
	return u.Repo.GetById(id)
}

func (u *UserUsecase) SignIn(id string) (domain.TokenResponse, error) {
	user, err := u.GetById(id)
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
	_, err := u.GetById(id)
	if err != nil {
		return err
	}

	return u.Repo.Update(id, updatedUser)
}

func (u *UserUsecase) ScanQR(id string) (domain.User, error) {
	user, err := u.GetById(id)
	if err != nil {
		return domain.User{}, err
	}

	now := time.Now()
	if user.LastEntered != nil && isSameDay(*user.LastEntered, now) {
		return user, domain.ErrUserAlreadyEntered
	}

	user.LastEntered = &now
	err = u.Update(id, &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserUsecase) UpdateRole(id string, role domain.Role) error {
	user, err := u.GetById(id)
	if err != nil {
		return err
	}
	user.Role = role
	return u.Update(id, &user)
}

func (u *UserUsecase) GetQRURL(id string) (string, error) {
	user, err := u.GetById(id)
	if err != nil {
		return "", err
	}

	baseURL := utils.GetEnv("PRODUCTION_BASE_URL", "http://localhost:4000/")
	return fmt.Sprintf("%s/api/users/qr/%s", baseURL, user.ID), nil
}

func (u *UserUsecase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func (u *UserUsecase) GetCardID(id string) (string, error) {
	user, err := u.GetById(id)
	if err != nil {
		return "", err
	}
	return user.ImageURL, nil
}

func (u *UserUsecase) AddStaff(phone string) error {
	user, err := u.Repo.GetByPhone(phone)
	if err != nil {
		return err
	}

	if user.Role == domain.Staff {
		return domain.ErrUserAlreadyStaff
	}

	user.Role = domain.Staff
	return u.Update(user.ID, &user)
}
