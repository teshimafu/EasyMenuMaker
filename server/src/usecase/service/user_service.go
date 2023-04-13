package service

import (
	"errors"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/service"
)

type UserService struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewUserService(userService *service.UserService, authService *service.AuthService) *UserService {
	return &UserService{
		userService: userService,
		authService: authService,
	}
}

func (u *UserService) GetUser(id string) (*entity.User, error) {
	return u.userService.GetUser(id)
}

func (u *UserService) GetUsers() ([]*entity.User, error) {
	return u.userService.GetAllUsers()
}

func (u *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user, err := u.userService.GetUserByAuth(email, password)
	if user != nil {
		return nil, errors.New("user is already exists")
	} else if err != nil {
		return nil, errors.New("email is duplicated")
	}
	return u.userService.CreateUser(name, email, password)
}

func (u *UserService) SignIn(email, password string) (string, error) {
	user, err := u.userService.GetUserByAuth(email, password)
	if err != nil {
		return "", err
	}
	token, err := u.authService.GenerateJWT(user.ID().Value())
	if err != nil {
		return "", err
	}
	return token.Value(), nil
}
