package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type ctxKey int

const Key ctxKey = 1

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private Key or Public Key cannot be nil")
	}
	return &Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil

}

func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {

	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenStr, err := tkn.SignedString(a.privateKey)
	if err != nil {
		return " ", fmt.Errorf("error in signing token %w", err)
	}

	return tokenStr, nil

}

func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims
	tkn, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})

	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("error while parsing claim %w", err)
	}

	if !tkn.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid Token")
	}
	return c, nil
}
