package converter

import (
	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/usecase/view"
)

type UserConverter struct{}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (up *UserConverter) ResponseUser(user *entity.User) *view.User {
	return up.ToJson(user)
}

func (up *UserConverter) ResponseUsers(users []*entity.User) []*view.User {
	usersJSON := make([]*view.User, len(users))
	for i, user := range users {
		user := up.ToJson(user)
		usersJSON[i] = user
	}
	return usersJSON
}

func (up *UserConverter) ToJson(user *entity.User) *view.User {
	if user == nil {
		return nil
	}
	id := user.ID().Value()
	userJSON := &view.User{
		ID:    &id,
		Name:  user.Name().Value(),
		Email: user.Email().Value(),
	}
	return userJSON
}
