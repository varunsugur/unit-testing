package repository

import (
	"context"
	"errors"
	"golang/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreatUser(ctx context.Context, userDetails models.User) (models.User, error) {

	result := r.Db.Create(&userDetails)
	if result.Error != nil {
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
