package service

import (
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

func (s *UserService) CreateUser(name, email string) (*entity.User, error) {
	return s.userService.CreateUser(name, email)
}
