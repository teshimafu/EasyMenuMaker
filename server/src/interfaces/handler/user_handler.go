package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"github.com/teshimafu/lazyPM/server/src/interfaces/presenter"
	"github.com/teshimafu/lazyPM/server/src/usecase/service"
)

type UserHandler struct {
	userService   *service.UserService
	userPresenter *presenter.UserPresenter
}

func NewUserHandler(userService *service.UserService, userPresenter *presenter.UserPresenter) *UserHandler {
	return &UserHandler{
		userService:   userService,
		userPresenter: userPresenter,
	}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	token := c.Get("token").(string)
	newToken, err := valueobject.NewToken(token)
	log.Println(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}
	users, err := h.userService.GetUsers(newToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return h.userPresenter.ResponseUsers(c, users)
}
