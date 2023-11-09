package service

import (
	"context"
	"fmt"
	"golang/internal/models"
	"strconv"
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
	fmt.Println("=========================", applicationData.Job.Experience)
	exp, err := strconv.Atoi(applicationData.Job.Experience)
	if err != nil {
		panic("error while parsing exp string from jobrequest")
	}
	fmt.Println("1111111111111")

	minexp, err := strconv.Atoi(val.MinExp)
	if err != nil {
		panic("error while parsing exp string from NewJob")
	}
	fmt.Println("22222222")

	if exp < minexp {
		return false, models.UserApplication{}, nil
	}

	if applicationData.Job.JobType != val.JobType {
		return false, models.UserApplication{}, nil
	}

	fmt.Println("3333333333333")

	np, err := strconv.Atoi(applicationData.Job.NoticePeriod)
	if err != nil {
		panic("error while parsing  np string from jobrequest")
	}
	fmt.Println("4444444444444")

	minnp, err := strconv.Atoi(val.MinNoticePeriod)
	if err != nil {
		panic("error while parsing  np string from NewJob")
	}
	fmt.Println("555555555555")

	if np < minnp {
		return false, models.UserApplication{}, nil
	}
	count := 0
	for _, v := range applicationData.Job.Location {
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
