package service

import (
	"context"
	"golang/internal/models"
)

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Job, error) {
	jobData, err := s.UserRepo.ViewJobDetailsBy(ctx, jid)
	if err != nil {
		return models.Job{}, nil
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Job, error) {
	jobData, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobData, nil

}

func (s *Service) AddJobDetails(ctx context.Context, jobData models.Job, cid uint64) (models.Job, error) {
	jobData.Cid = uint(cid)
	jobData, err := s.UserRepo.CreateJob(ctx, jobData)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil
}

func (s *Service) ViewJob(ctx context.Context, cid uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.FindJob(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
