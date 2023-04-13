package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

const (
	SecretKey = "your-secret-key"
)

type JWTGenerator struct {
}

func NewJWTGenerator() repository.ITokenRepository {
	return &JWTGenerator{}
}

func (j *JWTGenerator) GenerateToken(userID *value.UserID) (*value.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	return value.NewToken(tokenString)
}
