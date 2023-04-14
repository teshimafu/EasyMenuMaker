package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/entity"
)

type User struct {
	ID    *string `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
}
type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (up *UserPresenter) ResponseCreatedUser(c echo.Context, user *entity.User) error {
	if user == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusCreated, up.ToJson(user))
}

func (up *UserPresenter) ResponseUser(c echo.Context, user *entity.User) error {
	if user == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, up.ToJson(user))
}

func (up *UserPresenter) ResponseUsers(c echo.Context, users []*entity.User) error {
	usersJSON := make([]*User, len(users))
	for i, user := range users {
		user := up.ToJson(user)
		usersJSON[i] = user
	}

	return c.JSON(http.StatusOK, usersJSON)
}

func (up *UserPresenter) ToJson(user *entity.User) *User {
	if user == nil {
		return nil
	}
	id := user.ID().Value()
	userJSON := &User{
		ID:    &id,
		Name:  user.Name().Value(),
		Email: user.Email().Value(),
	}
	return userJSON
}
