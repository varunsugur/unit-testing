package middlewares

import (
	"errors"
	"golang/internal/auth"
)

type Mid struct {
	a *auth.Auth
}

func NewMid(a *auth.Auth) (Mid, error) {
	if a == nil {
		// log.Panic().Msg("Auth cant be nil")
		return Mid{}, errors.New("auth cant be nil")
	}
	return Mid{a: a}, nil
}
