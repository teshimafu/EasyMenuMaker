package valueobject

import (
	"errors"
	"log"
)

type UserID struct {
	value string
}

func NewUserID(value string) (*UserID, error) {
	if value == "" {
		log.Fatalln("user id is broken")
		return nil, errors.New("user id is broken")
	}

	return &UserID{value: value}, nil
}

func (n *UserID) Value() string {
	return n.value
}

func (n *UserID) Buffer() []byte {
	return []byte(n.value)
}

func (n *UserID) Equals(other *UserID) bool {
	return n.value == other.value
}
