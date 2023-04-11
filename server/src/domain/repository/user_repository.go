package repository

import (
	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type IUserRepository interface {
	FindByID(id *value.UserID) (*entity.User, error)
	FindByEmail(email *value.Email) (*entity.User, error)
	Find(input *entity.User) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Create(*entity.User) (*entity.User, error)
}
