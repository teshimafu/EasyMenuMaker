package service

import (
	"github.com/teshimafu/lazyPM/src/domain/entity"
	"github.com/teshimafu/lazyPM/src/domain/factory"
	"github.com/teshimafu/lazyPM/src/domain/repository"
	"github.com/teshimafu/lazyPM/src/domain/valueobject"
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
	return us.repo.Create(user)
}

func (s *UserService) GetUser(id *valueobject.UserID) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
