package entity

import value "github.com/teshimafu/lazyPM/src/domain/valueobject"

type User struct {
	id    *value.UserID
	name  *value.UserName
	email *value.Email
}

func NewUser(id, name, email string) (*User, error) {
	userID, err := value.NewUserID(id)
	if err != nil {
		return nil, err
	}
	userName, err := value.NewUserName(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := value.NewEmail(email)
	if err != nil {
		return nil, err
	}
	return &User{
		id:    userID,
		name:  userName,
		email: userEmail,
	}, nil
}

func (u *User) ID() *value.UserID {
	return u.id
}

func (u *User) Name() *value.UserName {
	return u.name
}

func (u *User) Email() *value.Email {
	return u.email
}
