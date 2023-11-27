package repository

import (
	"context"
	"errors"
	"fmt"
	"golang/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreatUser(ctx context.Context, userDetails models.User) (models.User, error) {

	result := r.Db.Create(&userDetails)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("error in creating user")
		return models.User{}, result.Error
	}
	return userDetails, nil
}

func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	var userDetail models.User
	result := r.Db.Where("email =?", email).First(&userDetail)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetail, nil
}

func (r *Repo) VerifyUser(vu models.VerifyUser) (models.User, error) {
	var userDetails models.User

	details := r.Db.Where("email=?", vu.UserEmail).First(&userDetails)
	if details.Error != nil {
		log.Info().Err(details.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil

}

func (r *Repo) ResetPassword(email string, reset string) error {
	result := r.Db.Model(&models.User{}).Where("email=?", email).Update("PasswordHash", reset)
	if result.Error != nil {
		fmt.Println("error while updating password", result.Error)
	} else {
		fmt.Println("password is reset successfully")
	}
	return nil
}
