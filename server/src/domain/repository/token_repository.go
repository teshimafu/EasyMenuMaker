package repository

import (
	value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"
)

type ITokenRepository interface {
	GenerateToken(userID *value.UserID) (*value.Token, error)
}
