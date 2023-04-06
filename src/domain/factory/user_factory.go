package factory

import (
	"github.com/google/uuid"
	"github.com/teshimafu/lazyPM/src/domain/entity"
)

type UserFactory interface {
	CreateUser(name, email string) (*entity.User, error)
}

type userFactory struct{}

func NewUserFactory() UserFactory {
	return &userFactory{}
}

func (u *userFactory) CreateUser(name, email string) (*entity.User, error) {
	id := uuid.New().String()
	return entity.NewUser(id, name, email)
}
