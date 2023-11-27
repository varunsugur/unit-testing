package repository

import (
	"context"
	"errors"
	"golang/internal/models"

	// "golang/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=repository

type UserRepo interface {
	CreatUser(ctx context.Context, userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)
	CreatCompany(ctx context.Context, data models.Company) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error)
	CreateJob(ctx context.Context, jobData models.Job) (models.ResponseJob, error)
	FindJob(ctx context.Context, cid uint64) ([]models.Job, error)
	FindAllJobs(ctx context.Context) ([]models.Job, error)
	ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Job, error)

	GetTheJobData(jobid uint) (models.Job, error)

	VerifyUser(vu models.VerifyUser) (models.User, error)
	ResetPassword(email string, resetpassword string) error
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		Db: db,
	}, nil
}
