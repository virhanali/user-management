package repository

import (
	"github.com/virhanali/user-management/domain/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(id int) (models.User, error)
	Update(user models.User) (models.User, error)
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

func (repo UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := repo.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo UserRepository) FindById(id int) (models.User, error) {
	var user models.User
	if err := repo.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo UserRepository) Update(user models.User) (models.User, error) {
	if err := repo.db.Debug().Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
