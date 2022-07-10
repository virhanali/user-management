package services

import (
	"github.com/jinzhu/copier"
	"github.com/virhanali/user-management/config"
	"github.com/virhanali/user-management/domain/models"
	"github.com/virhanali/user-management/repository"
)

type IUserService interface {
	CreateUser(user models.CreateUserRequest) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(userRequest models.UpdateUserRequest, id int) (models.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (service UserService) CreateUser(input models.CreateUserRequest) (models.UserResponse, error) {

	user := models.User{}
	copier.Copy(&user, &input)

	passwordHash, err := config.HashPassword(user.Password)

	if err != nil {
		return models.UserResponse{}, err
	}
	user.Password = passwordHash

	user, err = service.userRepository.Store(user)
	if err != nil {
		return models.UserResponse{}, err
	}
	userRes := models.UserResponse{}
	copier.Copy(&userRes, &user)

	return userRes, nil
}

func (service UserService) GetAllUser() ([]models.UserResponse, error) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userRes := []models.UserResponse{}
	copier.Copy(&userRes, &users)
	return userRes, nil
}

func (service UserService) GetUserById(id int) (models.UserResponse, error) {
	user, err := service.userRepository.FindById(id)
	if err != nil {
		return models.UserResponse{}, err
	}
	userRes := models.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil
}

func (services UserService) UpdateUser(userRequest models.UpdateUserRequest, id int) (models.UserResponse, error) {
	user, err := services.userRepository.FindById(id)

	if err != nil {
		return models.UserResponse{}, err
	}

	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true})

	if userRequest.Password != "" {
		passwordHash, err := config.HashPassword(user.Password)
		if err != nil {
			return models.UserResponse{}, err
		}
		user.Password = passwordHash
	}

	user, err = services.userRepository.Update(user)

	userRes := models.UserResponse{}

	copier.Copy(&userRes, &user)

	return userRes, nil
}
