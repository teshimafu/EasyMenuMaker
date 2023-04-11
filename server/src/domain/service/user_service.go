package service

import (
	"fmt"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/factory"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	"github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type UserService struct {
	repo    repository.IUserRepository
	factory factory.UserFactory
}

func NewUserService(repo repository.IUserRepository, factory factory.UserFactory) *UserService {
	return &UserService{
		repo:    repo,
		factory: factory,
	}
}

func (us *UserService) CreateUser(name, email string) (*entity.User, error) {
	user, err := us.factory.CreateUser(name, email)
	if err != nil {
		return nil, err
	}
	if _, err := us.repo.Find(user); err == nil {
		return nil, fmt.Errorf("user already exists")
	}
	return us.repo.Create(user)
}

func (s *UserService) GetUser(id *valueobject.UserID) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	return s.repo.FindAll()
}
