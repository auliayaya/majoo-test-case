package repository

import (
	"aulia-majoo-test/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &userRepository{DB: DB}
}

func (r *userRepository) GetUserByUsername(username string) (user *domain.User, err error) {

	err = r.DB.Debug().Preload("Merchants").Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {

	err := r.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
