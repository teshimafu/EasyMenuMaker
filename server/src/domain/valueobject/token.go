package valueobject

import "errors"

type Token struct {
	value string
}

func NewToken(value string) (*Token, error) {
	if value == "" {
		return nil, errors.New("token is required")
	}

	return &Token{value: value}, nil
}

func (n *Token) Value() string {
	return n.value
}

func (n *Token) Equals(other *Token) bool {
	return n.value == other.value
}
