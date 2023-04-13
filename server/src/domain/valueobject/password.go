package valueobject

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 15
)

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	return &Password{value: password}, nil
}

func HashPassword(password string) (*Password, error) {
	if password == "" {
		return nil, errors.New("password is required")
	} else if len(password) < 8 {
		return nil, errors.New("password is too short. password must be at least 8 characters")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &Password{value: string(hashedPassword)}, nil
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) Hash() []byte {
	return []byte(p.value)
}

func (p *Password) ComparePassword(rawPassword *Password) error {
	return bcrypt.CompareHashAndPassword(p.Hash(), rawPassword.Hash())
}
