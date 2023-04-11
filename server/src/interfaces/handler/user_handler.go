package handler

import (
	"fmt"
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

func (h *UserHandler) GetUser(c echo.Context) error {
	userID, err := valueobject.NewUserID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, err := h.userService.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
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

func (h *UserHandler) PostUser(c echo.Context) error {
	user := &presenter.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	fmt.Println(user)
	createdUser, err := h.userService.CreateUser(user.Name, user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return h.userPresenter.ResponseUser(c, createdUser)
}
