package service

import (
	"errors"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/service"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type UserService struct {
	userService *service.UserService
}

func NewUserService(userService *service.UserService) *UserService {
	return &UserService{
		userService: userService,
	}
}

func (s *UserService) GetUser(id *value.UserID) (*entity.User, error) {
	return s.userService.GetUser(id)
}

func (s *UserService) GetUsers() ([]*entity.User, error) {
	return s.userService.GetAllUsers()
}

func (s *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	emailValue, err := value.NewEmail(email)
	if err != nil {
		return nil, err
	}
	passwordValue, err := value.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := s.userService.GetUserByAuth(emailValue, passwordValue)
	if user != nil {
		return nil, errors.New("user is already exists")
	} else if err != nil {
		return nil, errors.New("email is duplicated")
	}
	return s.userService.CreateUser(name, email, password)
}
