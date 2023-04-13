package valueobject

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash []byte
}

func NewPassword(password []byte) (*Password, error) {
	return &Password{hash: password}, nil
}

func HashPassword(password string) (*Password, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Password{hash: hashedPassword}, nil
}

func (p *Password) Hash() []byte {
	return p.hash
}

func (p *Password) ComparePassword(rawPassword *Password) error {
	return bcrypt.CompareHashAndPassword(p.Hash(), rawPassword.Hash())
}
