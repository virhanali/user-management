package services

import (
	"github.com/jinzhu/copier"
	"github.com/virhanali/user-management/domain/models"
	"github.com/virhanali/user-management/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Store(user models.CreateUserRequest) (models.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (service UserService) Store(input models.CreateUserRequest) (models.UserResponse, error) {

	user := models.User{}
	copier.Copy(&user, &input)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return models.UserResponse{}, err
	}
	user.Password = string(passwordHash)

	user, err = service.userRepository.Store(user)
	if err != nil {
		return models.UserResponse{}, err
	}
	userRes := models.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, nil
}
