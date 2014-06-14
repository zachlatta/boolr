package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// NewToken creates a new Token from a provided user.
func NewToken(user *User) (*Token, error) {
	expires := time.Now().Add(time.Hour * 72)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["id"] = user.ID
	token.Claims["exp"] = expires.Unix()

	// TODO: Sign the token with an actual secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &Token{tokenString, expires}, nil
}
