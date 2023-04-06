package repository

import (
	"github.com/teshimafu/lazyPM/src/domain/entity"
	value "github.com/teshimafu/lazyPM/src/domain/valueobject"
)

type IUserRepository interface {
	FindByID(id *value.UserID) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Create(*entity.User) (*entity.User, error)
}
