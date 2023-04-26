package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/usecase/view"
)

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (up *UserPresenter) ResponseCreatedUser(c echo.Context, user *view.User) error {
	if user == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusCreated, user)
}

func (up *UserPresenter) ResponseUser(c echo.Context, user *view.User) error {
	if user == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (up *UserPresenter) ResponseUsers(c echo.Context, users []*view.User) error {

	return c.JSON(http.StatusOK, users)
}
