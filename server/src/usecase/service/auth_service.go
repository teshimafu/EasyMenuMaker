package service

import (
	"github.com/teshimafu/lazyPM/server/src/domain/factory"
	"github.com/teshimafu/lazyPM/server/src/domain/service"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"github.com/teshimafu/lazyPM/server/src/usecase/converter"
	"github.com/teshimafu/lazyPM/server/src/usecase/view"
)

type AuthService struct {
	userService   *service.UserService
	authService   *service.AuthService
	userFactory   factory.UserFactory
	userConverter *converter.UserConverter
}

func NewAuthService(userService *service.UserService, authService *service.AuthService, userFactory factory.UserFactory, userConverter *converter.UserConverter) *AuthService {
	return &AuthService{
		userService:   userService,
		authService:   authService,
		userFactory:   userFactory,
		userConverter: userConverter,
	}
}

func (u *AuthService) GetUserID(token *value.Token) (*value.UserID, error) {
	return u.authService.GetUserID(token)
}

func (u *AuthService) Signup(form *view.SignupForm) (*view.User, error) {
	user, err := u.userFactory.CreateUser(form.Name, form.Email, form.Password)
	if err != nil {
		return nil, err
	}
	user, err = u.userService.Signup(user)
	if err != nil {
		return nil, err
	}
	return u.userConverter.ResponseUser(user), nil
}

func (u *AuthService) Signin(email *value.Email, password *value.Password) (string, error) {
	user, err := u.userService.GetUserByAuth(email, password)
	if err != nil {
		return "", err
	}
	token, err := u.authService.GenerateToken(user.ID())
	if err != nil {
		return "", err
	}
	return token.Value(), nil
}
