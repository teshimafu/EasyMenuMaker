package valueobject

import "errors"

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("user name is required")
	}

	return &UserName{value: value}, nil
}

func (n *UserName) Value() string {
	return n.value
}

func (n *UserName) Equals(other *UserID) bool {
	return n.value == other.value
}
