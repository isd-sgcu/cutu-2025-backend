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
	return u.Repo.Create(user)
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.Repo.GetAll()
}
