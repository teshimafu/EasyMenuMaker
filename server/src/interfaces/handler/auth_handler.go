package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/interfaces/presenter"
	"github.com/teshimafu/lazyPM/server/src/usecase/service"
)

type SignupForm struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthHandler struct {
	userService   *service.UserService
	userPresenter *presenter.UserPresenter
}

func NewAuthHandler(userService *service.UserService, userPresenter *presenter.UserPresenter) *AuthHandler {
	return &AuthHandler{
		userService:   userService,
		userPresenter: userPresenter,
	}
}

func (a *AuthHandler) PostSignup(c echo.Context) error {
	user := &SignupForm{}
	if err := c.Bind(user); err != nil {
		return err
	}

	createdUser, err := a.userService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return a.userPresenter.ResponseUser(c, createdUser)
}
