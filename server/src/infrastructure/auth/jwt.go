package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type JWTGenerator struct {
}

func NewJWTGenerator() repository.ITokenRepository {
	return &JWTGenerator{}
}

func (j *JWTGenerator) GenerateToken(userID *value.UserID) (*value.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return value.NewToken(tokenString)
}
