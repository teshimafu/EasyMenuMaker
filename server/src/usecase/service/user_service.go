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

func (u *UserService) GetUser(id *value.UserID) (*entity.User, error) {
	return u.userService.GetUser(id)
}

func (u *UserService) GetUsers() ([]*entity.User, error) {
	return u.userService.GetAllUsers()
}
