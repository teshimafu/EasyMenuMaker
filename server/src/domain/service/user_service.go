package service

import (
	"errors"

	"github.com/teshimafu/lazyPM/server/src/domain/entity"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
	"gorm.io/gorm"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserNotFound    = errors.New("user is not found")
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Signup(user *entity.User) (*entity.User, error) {
	return us.repo.Create(user)
}

func (s *UserService) GetUser(id *value.UserID) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUserByAuth(email *value.Email, password *value.Password) (*entity.User, error) {
	user, err := s.repo.FindByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	if err := user.Password().ComparePassword(password); err != nil {
		return nil, ErrInvalidPassword
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*entity.User, error) {
	return s.repo.FindAll()
}
