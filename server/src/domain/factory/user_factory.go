package factory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"gorm.io/gorm"
)

type UserFactory interface {
	CreateUser(name, email, password string) (*entity.User, error)
}

type userFactory struct {
	repo repository.IUserRepository
}

func NewUserFactory(repo repository.IUserRepository) UserFactory {
	return &userFactory{
		repo: repo,
	}
}

func (u *userFactory) CreateUser(name, email, password string) (*entity.User, error) {
	userEmail, err := value.NewEmail(email)
	if err != nil {
		return nil, err
	}
	if _, err := u.repo.FindByEmail(userEmail); !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user is already exists")
	}
	id := uuid.New().String()
	hashPassword, err := value.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(id, name, email, hashPassword.Value())
}
