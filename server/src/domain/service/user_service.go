package service

import (
	"errors"
	"fmt"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/factory"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	"github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"gorm.io/gorm"
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

func (us *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user, err := us.factory.CreateUser(name, email, password)
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

func (s *UserService) GetUserByAuth(email *valueobject.Email, password *valueobject.Password) (*entity.User, error) {
	user, err := s.repo.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	if err := password.ComparePassword(user.Password()); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	return s.repo.FindAll()
}
