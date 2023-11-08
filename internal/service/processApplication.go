package service

import (
	"context"
	"golang/internal/models"
	"sync"
)

func (s *Service) ProccessApplication(ctx context.Context, applicationData []models.UserApplication) ([]models.UserApplication, error) {
	var wg = new(sync.WaitGroup)
	var finalData []models.UserApplication
	for _, v := range applicationData {
		wg.Add(1)
		go func(v models.UserApplication) {
			defer wg.Done()
			check, v, err := s.compareAndCheck(v)
			if err != nil {
				return
			}
			if check {
				finalData = append(finalData, v)
			}
		}(v)

		// check, v, err := s.compareAndCheck(v)

		// if err != nil {
		// 	return nil, err
		// }
		// if check {
		// 	finalData = append(finalData, v)
		// }
	}
	wg.Wait()
	return finalData, nil
}

var cacheMap = make(map[uint]models.Job)

func (s *Service) compareAndCheck(applicationData models.UserApplication) (bool, models.UserApplication, error) {
	val, exists := cacheMap[applicationData.Jid]
	if !exists {
		jobData, err := s.UserRepo.GetTheJobData(applicationData.Jid)
		if err != nil {
			return false, models.UserApplication{}, err
		}
		cacheMap[applicationData.Jid] = jobData
		val = jobData
	}
	if applicationData.Job.MinExp < val.MinExp {
		return false, models.UserApplication{}, nil
	}
	if applicationData.Job.JobType != val.JobType {
		return false, models.UserApplication{}, nil
	}
	if applicationData.Job.MinNoticePeriod < val.MinNoticePeriod {
		return false, models.UserApplication{}, nil
	}
	count := 0
	for _, v := range applicationData.Job.JobLocation {
		for _, v1 := range val.JobLocation {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}
	count = 0
	for _, v := range applicationData.Job.Qualifications {
		for _, v1 := range val.Qualifications {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}
	count = 0
	for _, v := range applicationData.Job.Technology {
		for _, v1 := range val.Technology {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}
	count = 0
	for _, v := range applicationData.Job.Shift {
		for _, v1 := range val.Shift {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, models.UserApplication{}, nil
	}

	return true, applicationData, nil
}
