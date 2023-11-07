package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"fmt"
	"golang/internal/models"
	"golang/internal/pkg"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) UserSignup(ctx context.Context, nu models.NewUser) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, fmt.Errorf("generating Password hash %w", err)
	}

	u := models.User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPassword),
	}
	userDetails, err := s.UserRepo.CreatUser(ctx, u)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}

func (s *Service) UserSignin(ctx context.Context, user models.UserLogin) (string, error) {
	var userDetails models.User
	userDetails, err := s.UserRepo.CheckEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	err = pkg.CheckHashedPassword(user.Password, userDetails.PasswordHash)
	if err != nil {
		log.Error().Err(err).Send()
		return "", errors.New("entered password is wrong")
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token, err := s.auth.GenerateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil

}
