package service

import (
	"context"
	"errors"
	"golang/internal/auth"
	"golang/internal/cache"
	"golang/internal/models"
	"golang/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     *auth.Auth
	rdb      cache.Cache
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service

type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserSignin(ctx context.Context, user models.UserLogin) (string, error)
	AddCompanyDetails(ctx context.Context, data models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyDetail(ctx context.Context, cid uint64) (models.Company, error)
	ViewJob(ctx context.Context, cid uint64) ([]models.Job, error)

	AddJobDetails(ctx context.Context, jobData models.NewJob, cid uint64) (models.ResponseJob, error)
	ViewAllJobs(ctx context.Context) ([]models.Job, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Job, error)

	ProccessApplication(ctx context.Context, applicationData []models.UserApplication) ([]models.UserApplication, error)

	VerifyUser(ctx context.Context, vu models.VerifyUser) error
}

func NewService(userRepo repository.UserRepo, a *auth.Auth, rdb cache.Cache) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
		rdb:      rdb,
	}, nil
}
