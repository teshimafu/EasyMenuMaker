package service

import (
	"github.com/teshimafu/lazyPM/server/src/domain/service"
	"github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"github.com/teshimafu/lazyPM/server/src/usecase/converter"
	"github.com/teshimafu/lazyPM/server/src/usecase/view"
)

type UserService struct {
	userService   *service.UserService
	userConverter *converter.UserConverter
}

func NewUserService(userService *service.UserService, userConverter *converter.UserConverter) *UserService {
	return &UserService{
		userService:   userService,
		userConverter: userConverter,
	}
}

func (u *UserService) GetUser(id string) (*view.User, error) {
	userID, err := valueobject.NewUserID(id)
	if err != nil {
		return nil, err
	}

	user, err := u.userService.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return u.userConverter.ResponseUser(user), nil
}

func (u *UserService) GetUsers() ([]*view.User, error) {
	user, err := u.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return u.userConverter.ResponseUsers(user), nil
}
