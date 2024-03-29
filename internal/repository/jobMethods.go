package repository

import (
	"context"
	"errors"
	"golang/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Job, error) {
	var jobData models.Job
	result := r.Db.Where("id = ?", jid).Find(&jobData)
	if result.Error != nil {
		log.Error().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create the jobs")
	}
	return jobData, nil
}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Job) (models.ResponseJob, error) {
	result := r.Db.Create(&jobData)
	if result.Error != nil {
		log.Error().Err(result.Error).Send()
		return models.ResponseJob{}, errors.New("could not create the jobs")
	}
	return models.ResponseJob{
		Id: jobData.ID,
	}, nil
}

func (r *Repo) FindAllJobs(ctx context.Context) ([]models.Job, error) {
	var jobDatas []models.Job
	result := r.Db.Find(&jobDatas)
	if result.Error != nil {
		log.Error().Err(result.Error)
		return nil, errors.New("could not find the jobs")
	}
	return jobDatas, nil

}

func (r *Repo) FindJob(ctx context.Context, cid uint64) ([]models.Job, error) {
	var jobData []models.Job
	result := r.Db.Where("cid = ?", cid).Find(&jobData)
	if result.Error != nil {
		log.Error().Err(result.Error).Send()
		return nil, errors.New("could not find the company")
	}
	return jobData, nil
}

func (r *Repo) GetTheJobData(jobid uint) (models.Job, error) {
	var jobData models.Job

	// Preload related data using GORM's Preload method
	result := r.Db.Preload("Company").
		Preload("JobLocation").
		Preload("Technology").
		Preload("Qualifications").
		Preload("Shift").
		Where("id = ?", jobid).
		Find(&jobData)

	if result.Error != nil {
		log.Error().Err(result.Error).Send()
		return models.Job{}, result.Error
	}

	return jobData, nil
}
