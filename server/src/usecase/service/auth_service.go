package service

import (
	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/service"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type AuthService struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewAuthService(userService *service.UserService, authService *service.AuthService) *AuthService {
	return &AuthService{
		userService: userService,
		authService: authService,
	}
}

func (u *AuthService) GetUserID(token *value.Token) (*value.UserID, error) {
	return u.authService.GetUserID(token)
}

func (u *AuthService) Signup(user *entity.User) (*entity.User, error) {
	return u.userService.Signup(user)
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
