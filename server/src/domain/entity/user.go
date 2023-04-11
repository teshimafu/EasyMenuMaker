package entity

import value "github.com/teshimafu/lazyPM/server/src/domain/valueobject"

type User struct {
	id       *value.UserID
	name     *value.UserName
	email    *value.Email
	password *value.Password
}

func NewUser(id, name, email string, password []byte) (*User, error) {
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
	userPassword, err := value.NewPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		id:       userID,
		name:     userName,
		email:    userEmail,
		password: userPassword,
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

func (u *User) Password() *value.Password {
	return u.password
}
