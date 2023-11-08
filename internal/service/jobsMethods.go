package service

import (
	"context"
	"golang/internal/models"

	"gorm.io/gorm"
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

func (s *Service) AddJobDetails(ctx context.Context, jobData models.NewJob, cid uint64) (models.ResponseJob, error) {
	jobData.Cid = uint(cid)

	creatJob := models.Job{
		Cid:             jobData.Cid,
		Budget:          jobData.Budget,
		MinNoticePeriod: jobData.MinNoticePeriod,
		MaxNoticePeriod: jobData.MaxNoticePeriod,

		Description: jobData.Description,
		MinExp:      jobData.MinExp,
		MaxExp:      jobData.MaxExp,

		JobType: jobData.JobType,
	}

	for _, v := range jobData.JobLocation {
		tempCreatJobDetails := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		creatJob.JobLocation = append(creatJob.JobLocation, tempCreatJobDetails)
	}

	for _, v := range jobData.Technology {
		tempCreatJobDeails := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		creatJob.Technology = append(creatJob.Technology, tempCreatJobDeails)
	}

	for _, v := range jobData.Qualifications {
		tempCreatJobDeails := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		creatJob.Qualifications = append(creatJob.Qualifications, tempCreatJobDeails)
	}

	for _, v := range jobData.Shift {
		tempCreatJobDeails := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		creatJob.Shift = append(creatJob.Shift, tempCreatJobDeails)
	}

	jobDetails, err := s.UserRepo.CreateJob(ctx, creatJob)
	if err != nil {
		return models.ResponseJob{}, err
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
