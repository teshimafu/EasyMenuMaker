package valueobject

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("email is required")
	}
	if !isValidEmail(value) {
		return nil, errors.New("invalid email format")
	}
	return &Email{value: value}, nil
}

func isValidEmail(value string) bool {
	validEmailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return validEmailPattern.MatchString(value)
}

func (n *Email) Value() string {
	return n.value
}

func (n *Email) Equals(other *Email) bool {
	return n.value == other.value
}
