package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user is not found"})
	}

	return h.userPresenter.ResponseUser(c, user)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return h.userPresenter.ResponseUsers(c, users)
}
