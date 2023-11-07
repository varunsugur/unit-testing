package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"golang/internal/models"
)

func (r *Repo) CreatCompany(ctx context.Context, data models.Company) (models.Company, error) {
	result := r.Db.Create(&data)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not create the company")

	}
	return data, nil
}

func (r *Repo) ViewCompanies(ctx context.Context) ([]models.Company, error) {
	var companiesData []models.Company
	result := r.Db.Find(&companiesData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the companies")
	}
	return companiesData, nil

}

func (r *Repo) ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error) {
	var companyData models.Company
	result := r.Db.Where("id = ?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not find the company")
	}
	return companyData, nil
}
