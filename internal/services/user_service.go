package services

import (
	dto "sme-demo/internal/dto/authentication"
	"sme-demo/internal/repositories/user"
)

type UserServiceInterface interface {
	Create(user dto.SignUpUserDto[any]) (*user.User[any], error)
}

// UserService is the implementation of the UserService interface
type UserService struct {
	userRepository *user.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepository *user.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Create a new user by a user. Will check if the user is allowed to create a new user.
func (s *UserService) Create(user dto.SignUpUserDto[any], byUser user.User[any]) (*user.User[any], error) {
	//TODO: check if the user is allowed to create a new user
	return s.userRepository.Create(user)
}
