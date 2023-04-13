package service

import (
	"log"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/service"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
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

func (u *UserService) GetUser(id *value.UserID) (*entity.User, error) {
	return u.userService.GetUser(id)
}

func (u *UserService) GetUsers(token *value.Token) ([]*entity.User, error) {
	// wip
	id, _ := u.authService.GetUserID(token)
	log.Println(id.Value())
	return u.userService.GetAllUsers()
}

func (u *UserService) Signup(user *entity.User) (*entity.User, error) {
	return u.userService.Signup(user)
}

func (u *UserService) Signin(email *value.Email, password *value.Password) (string, error) {
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
