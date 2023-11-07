package service

import (
	"context"
	"errors"
	"golang/internal/auth"
	"golang/internal/models"
	"golang/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     *auth.Auth
}

type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserSignin(ctx context.Context, user models.UserLogin) (string, error)
	AddCompanyDetails(ctx context.Context, data models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyDetail(ctx context.Context, cid uint64) (models.Company, error)
	ViewJob(ctx context.Context, cid uint64) ([]models.Job, error)

	AddJobDetails(ctx context.Context, jobData models.Job, cid uint64) (models.Job, error)
	ViewAllJobs(ctx context.Context) ([]models.Job, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Job, error)
}

func NewService(userRepo repository.UserRepo, a *auth.Auth) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
	}, nil
}
