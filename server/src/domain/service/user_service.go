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

var (
	ErrInvalidPassword = errors.New("invalid password")
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

func (s *UserService) GetUser(id string) (*entity.User, error) {
	userID, err := valueobject.NewUserID(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(userID)
}

func (s *UserService) GetUserByAuth(email, password string) (*entity.User, error) {
	userEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindByEmail(userEmail)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	userPassword, err := valueobject.NewPassword([]byte(password))
	if err != nil {
		return nil, err
	}
	if err := user.Password().ComparePassword(userPassword); err != nil {
		return nil, ErrInvalidPassword
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	return s.repo.FindAll()
}
