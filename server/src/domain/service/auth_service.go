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

func (s *AuthService) GenerateToken(userID *value.UserID) (*value.Token, error) {
	return s.repo.GenerateToken(userID)
}

func (s *AuthService) GetUserID(token *value.Token) (*value.UserID, error) {
	return s.repo.GetUserID(token)
}
