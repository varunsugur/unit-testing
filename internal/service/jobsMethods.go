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

func (s *Service) AddJobDetails(ctx context.Context, jobData models.NewJob, cid uint64) (models.Job, error) {
	jobData.Cid = uint(cid)

	creatJob := models.Job{
		Cid:             jobData.Cid,
		Budget:          jobData.Budget,
		MinNoticePeriod: jobData.MinNoticePeriod,
		MaxNoticePeriod: jobData.MaxNoticePeriod,
		JobLocation:     jobData.JobLocation,
		Technology:      jobData.Technology,
		Description:     jobData.Description,
		MinExp:          jobData.MinExp,
		MaxExp:          jobData.MaxExp,
		Qualifications:  jobData.Qualifications,
		Shift:           jobData.Shift,
		JobType:         jobData.JobType,
	}

	jobDetails, err := s.UserRepo.CreateJob(ctx, creatJob)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, nil
}

func (s *Service) ViewJob(ctx context.Context, cid uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.FindJob(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
