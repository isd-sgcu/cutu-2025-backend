package repository

import (
	"github.com/isd-sgcu/cutu2025-backend/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

// GetAll implements usecase.UserRepositoryInterface.
func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Find(&users).Error
	return users, err
}
