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

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetById(id string) (domain.User, error) {
	var user domain.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *UserRepository) Update(id string, user *domain.User) error {
	err := r.DB.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error
	return err
}

func (r *UserRepository) Delete(id string) error {
	err := r.DB.Where("id = ?", id).Delete(&domain.User{}).Error
	return err
}
