package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"github.com/teshimafu/lazyPM/server/src/interfaces/presenter"
	"github.com/teshimafu/lazyPM/server/src/usecase/service"
	"github.com/teshimafu/lazyPM/server/src/usecase/view"
)

type AuthHandler struct {
	authService   *service.AuthService
	userPresenter *presenter.UserPresenter
}

func NewAuthHandler(authService *service.AuthService, userPresenter *presenter.UserPresenter) *AuthHandler {
	return &AuthHandler{
		authService:   authService,
		userPresenter: userPresenter,
	}
}

func (a *AuthHandler) PostSignup(c echo.Context) error {
	userCmd := &view.SignupForm{}
	if err := c.Bind(userCmd); err != nil {
		return err
	}

	createdUser, err := a.authService.Signup(userCmd)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return a.userPresenter.ResponseCreatedUser(c, createdUser)
}

func (a *AuthHandler) PostSignin(c echo.Context) error {
	req := &view.SigninForm{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	email, err := valueobject.NewEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	password, err := valueobject.NewPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := a.authService.Signin(email, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
