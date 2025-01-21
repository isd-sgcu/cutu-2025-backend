package usecase

import "github.com/isd-sgcu/cutu2025-backend/domain"

type UserUsecase struct {
	Repo UserRepositoryInterface
}

type UserRepositoryInterface interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
}

func NewUserUsecase(repo UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{Repo: repo}
}

func (u *UserUsecase) Register(user *domain.User) error {
	u.assignRole(user)
	return u.Repo.Create(user)
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.Repo.GetAll()
}

// assing role based on phone number
func (u *UserUsecase) assignRole(user *domain.User) {
	// mock phone number
	phone_tel_list := []string{"06", "08", "09"}
	user.Role = domain.Student
	if user.Phone != "" {
		for _, tel := range phone_tel_list {
			if user.Phone == tel {
				user.Role = domain.Staff
				break
			}
		}
	}
}
