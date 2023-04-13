package service

import (
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type AuthService struct {
	repo repository.ITokenRepository
}

func NewAuthService(repo repository.ITokenRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) GenerateJWT(id string) (*value.Token, error) {
	userID, err := value.NewUserID(id)
	if err != nil {
		return nil, err
	}
	return s.repo.GenerateToken(userID)
}
