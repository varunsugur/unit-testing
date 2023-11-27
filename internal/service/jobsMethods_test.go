package service

import (
	"context"
	"errors"
	"golang/internal/auth"
	"golang/internal/cache"
	"golang/internal/models"
	"golang/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.NewJob
		cid     uint64
	}
	tests := []struct {
		name         string
		args         args
		want         models.ResponseJob
		wantErr      bool
		mockResponse func() (models.ResponseJob, error)
	}{
		// TODO: Add test cases.
		{
			name: "error from repository layer",
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					Name:            "Software Developer",
					Budget:          "5lack",
					MinNoticePeriod: "7",
					MaxNoticePeriod: "30",
					JobLocation: []uint{
						uint(1), uint(2),
					},
					Technology: []uint{
						uint(1), uint(4),
					},
					Description: "Develop-Software Application and Code-Review",
					MinExp:      "0",
					MaxExp:      "3",
					Qualifications: []uint{
						uint(3), uint(4),
					},
					Shift: []uint{
						uint(1),
					},
					JobType: "Full-Time",
				},
				cid: 1,
			},
			want:    models.ResponseJob{},
			wantErr: true,
			mockResponse: func() (models.ResponseJob, error) {
				return models.ResponseJob{}, errors.New("test error from mock function")
			},
		},
		{
			name: "success from repository layer",
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					Name:            "Software Developer",
					Budget:          "5lack",
					MinNoticePeriod: "7",
					MaxNoticePeriod: "30",
					JobLocation: []uint{
						uint(1), uint(2),
					},
					Technology: []uint{
						uint(1), uint(4),
					},
					Description: "Develop-Software Application and Code-Review",
					MinExp:      "0",
					MaxExp:      "3",
					Qualifications: []uint{
						uint(3), uint(4),
					},
					Shift: []uint{
						uint(1),
					},
					JobType: "Full-Time",
				},
				cid: 1,
			},
			want: models.ResponseJob{
				Id: 1,
			},
			wantErr: false,
			mockResponse: func() (models.ResponseJob, error) {
				return models.ResponseJob{
					Id: 1,
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mr := repository.NewMockUserRepo(mc)

			mr.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			svc, err := NewService(mr, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := svc.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllJobs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		args         args
		want         []models.Job
		wantErr      bool
		mockResponse func() ([]models.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "database success",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Job{
				// {
				// 	Cid:          1,
				// 	Name:         "Software Engineer",
				// 	Salary:       "20000",
				// 	NoticePeriod: "3months",
				// },
				// {
				// 	Cid:          1,
				// 	Name:         "Softaware Trainee",
				// 	Salary:       "20000",
				// 	NoticePeriod: "3 months",
				// },
			},
			wantErr: false,
			mockResponse: func() ([]models.Job, error) {
				return []models.Job{
					// {
					// 	Cid:          1,
					// 	Name:         "Software Engineer",
					// 	Salary:       "20000",
					// 	NoticePeriod: "3 months",
					// },
					// {
					// 	Cid:          1,
					// 	Name:         "Software Trainee",
					// 	Salary:       "20000",
					// 	NoticePeriod: "3 months",
					// },
				}, nil
			},
		},
		{
			name: "error from databse",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockResponse: func() ([]models.Job, error) {
				return nil, errors.New("could not find jobs in databse")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mr := repository.NewMockUserRepo(mc)

			mr.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			svc, err := NewService(mr, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := svc.ViewAllJobs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewJobById(t *testing.T) {
	type args struct {
		ctx context.Context
		jid uint64
	}
	tests := []struct {
		name             string
		s                *Service
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		// TODO: Add test cases.
		// {
		// 	name: "successful",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		jid: 6,
		// 	},
		// 	want: models.Job{
		// 		Company: models.Company{
		// 			Name: "Info",
		// 		},
		// 		Cid:          1,
		// 		Name:         "Software Developer",
		// 		Salary:       "20000",
		// 		NoticePeriod: "3 months",
		// 	},
		// 	wantErr: false,
		// 	mockresponse: func() (models.Job, error) {
		// 		return models.Job{
		// 			Company: models.Company{
		// 				Name: "Info",
		// 			},
		// 			Cid:          1,
		// 			Name:         "Software Developer",
		// 			Salary:       "20000",
		// 			NoticePeriod: "3 months",
		// 		}, nil
		// 	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, tt.args.jid).Return(tt.mockRepoResponse()).AnyTimes()
			s, err := NewService(mockRepo, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			got, err := s.ViewJobById(tt.args.ctx, tt.args.jid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewJob(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name         string
		args         args
		want         []models.Job
		wantErr      bool
		mockResponse func() ([]models.Job, error)
	}{
		// TODO: Add test cases.
		// {
		// 	name: "success",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		cid: 1,
		// 	},
		// 	want: []models.Job{
		// 		{
		// 			Cid:          1,
		// 			Name:         "software Engineer",
		// 			Salary:       "20000",
		// 			NoticePeriod: "3 months",
		// 		},
		// 		{
		// 			Cid:          1,
		// 			Name:         "Software Trainee",
		// 			Salary:       "20000",
		// 			NoticePeriod: "3 months",
		// 		},
		// 	},
		// 	wantErr: false,
		// 	mockResponse: func() ([]models.Job, error) {
		// 		return []models.Job{
		// 			{
		// 				Cid:          1,
		// 				Name:         "software Engineer",
		// 				Salary:       "20000",
		// 				NoticePeriod: "3 months",
		// 			},
		// 			{
		// 				Cid:          1,
		// 				Name:         "Software Trainee",
		// 				Salary:       "20000",
		// 				NoticePeriod: "3 months",
		// 			},
		// 		}, nil
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().FindJob(tt.args.ctx, tt.args.cid).Return(tt.mockResponse()).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("could not initialize service layer")
				return
			}
			got, err := s.ViewJob(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
