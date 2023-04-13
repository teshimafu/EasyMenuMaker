package service

import (
	"errors"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/service"
)

type UserService struct {
	userService *service.UserService
}

func NewUserService(userService *service.UserService) *UserService {
	return &UserService{
		userService: userService,
	}
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	return s.userService.GetUser(id)
}

func (s *UserService) GetUsers() ([]*entity.User, error) {
	return s.userService.GetAllUsers()
}

func (s *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user, err := s.userService.GetUserByAuth(email, password)
	if user != nil {
		return nil, errors.New("user is already exists")
	} else if err != nil {
		return nil, errors.New("email is duplicated")
	}
	return s.userService.CreateUser(name, email, password)
}
