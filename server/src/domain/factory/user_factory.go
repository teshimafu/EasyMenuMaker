package factory

import (
	"github.com/google/uuid"
	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type UserFactory interface {
	CreateUser(name, email, password string) (*entity.User, error)
}

type userFactory struct{}

func NewUserFactory() UserFactory {
	return &userFactory{}
}

func (u *userFactory) CreateUser(name, email, password string) (*entity.User, error) {
	id := uuid.New().String()
	hashPassword, err := valueobject.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(id, name, email, hashPassword.Hash())
}
