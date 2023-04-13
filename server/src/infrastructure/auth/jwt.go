package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/teshimafu/lazyPM/server/src/domain/repository"
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type JWTGenerator struct {
	secretKey string
}

func NewJWTGenerator() repository.ITokenRepository {
	return &JWTGenerator{
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (j *JWTGenerator) GenerateToken(userID *value.UserID) (*value.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID.Value()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	return value.NewToken(tokenString)
}

func (j *JWTGenerator) GetUserID(token *value.Token) (*value.UserID, error) {
	jwtToken, err := jwt.Parse(token.Value(), func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, fmt.Errorf("user_id claim not found or not a string")
		}
		return value.NewUserID(userID)
	}

	return nil, fmt.Errorf("invalid token")
}
