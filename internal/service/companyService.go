package service

import (
	"context"
	"golang/internal/models"
)

func (s *Service) AddCompanyDetails(ctx context.Context, data models.Company) (models.Company, error) {
	companyData, err := s.UserRepo.CreatCompany(ctx, data)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil

}

func (s *Service) ViewAllCompanies(ctx context.Context) ([]models.Company, error) {

	companies, err := s.UserRepo.ViewCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *Service) ViewCompanyDetail(ctx context.Context, cid uint64) (models.Company, error) {
	companydata, err := s.UserRepo.ViewCompanyById(ctx, cid)
	if err != nil {
		return models.Company{}, err
	}
	return companydata, nil
}
