package repository

import (
	"github.com/virhanali/user-management/domain/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user models.User) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) Store(user models.User) (models.User, error) {
	if err := repo.db.Debug().Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
