package valueobject

import "errors"

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("email is required")
	}

	return &Email{value: value}, nil
}

func (n *Email) Value() string {
	return n.value
}

func (n *Email) Equals(other *Email) bool {
	return n.value == other.value
}
