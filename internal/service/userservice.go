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
	// "github.com/rs/zerolog/log"
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
		DOB:          nu.DOB,
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

func (s *Service) VerifyUser(ctx context.Context, vu models.VerifyUser) error {
	details, err := s.UserRepo.VerifyUser(vu)
	if err != nil {
		return errors.New("email is not found")

	}

	if vu.DOB != details.DOB {
		return errors.New("enter correct date of birth")
	}
	otp, err := pkg.GenerateOTP(vu.UserEmail)
	if err != nil {
		return errors.New("could not send otp to verified email")
	}

	err = s.rdb.AddtoOTPCache(ctx, vu.UserEmail, otp)
	if err != nil {
		return errors.New("failed to store in cache")
	}

	return nil
}

func (s *Service) UpdatePassword(ctx context.Context, details models.ResetDetails) error {

	if details.NewPassword != details.ConfirmPassword {
		return errors.New("newpassword and confirm password are not same")
	}

	otp, err := s.rdb.GetCacheOtp(ctx, details.Email)
	if err != nil {
		return err
	}

	if otp != details.OTP {
		return errors.New("otp mismatched")
	}

	// err = s.rdb.DeleteCacheOtp(ctx, details.Email)
	// if err != nil {
	// 	return errors.New()
	// }

	hashPassword, err := pkg.HashPassword(details.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	s.UserRepo.ResetPassword(details.Email, hashPassword)
	// g
	return nil

}
