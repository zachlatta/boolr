package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	ErrInvalidUserUsername = errors.New("invalid username")
	ErrInvalidUserPassword = errors.New("invalid password")
)

type User struct {
	ID       int64     `json:"id"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Username string    `json:"username"`
	Password string    `json:"-"`
}

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(jsonReader io.Reader) (*User, error) {
	var rU RequestUser
	if err := json.NewDecoder(jsonReader).Decode(&rU); err != nil {
		return nil, err
	}

	if err := rU.validate(); err != nil {
		return nil, err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(rU.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Username: rU.Username,
		Password: string(b),
	}

	return &user, nil
}

func (u *RequestUser) validate() error {
	switch {
	case len(u.Username) >= 255 || len(u.Username) <= 0:
		return ErrInvalidUserUsername
	case len(u.Password) < 6:
		return ErrInvalidUserPassword
	case len(u.Password) > 256:
		return ErrInvalidUserPassword
	default:
		return nil
	}
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
