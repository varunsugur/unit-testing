package service

// import (
// 	"context"
// 	"errors"
// 	"golang/internal/auth"
// 	"golang/internal/models"
// 	"golang/internal/repository"
// 	"reflect"
// 	"testing"

// 	"go.uber.org/mock/gomock"
// )

// func TestService_AddJobDetails(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		jobData models.Job
// 		cid     uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		want         models.Job
// 		wantErr      bool
// 		mockResponse func() (models.Job, error)
// 	}{
// TODO: Add test cases.
// {
// 	name: "successful",
// 	args: args{
// 		ctx: context.Background(),
// 		jobData: models.Job{
// 			Name:         "Software Engineer",
// 			Salary:       "20000",
// 			NoticePeriod: "3 months",
// 		},
// 		cid: 1,
// 	},
// 	want: models.Job{
// 		Cid:          1,
// 		Name:         "Software Engineer",
// 		Salary:       "20000",
// 		NoticePeriod: "3 months",
// 	},
// 	wantErr: false,
// 	mockResponse: func() (models.Job, error) {
// 		return models.Job{
// 			Cid:          1,
// 			Name:         "Software Engineer",
// 			Salary:       "20000",
// 			NoticePeriod: "3 months",
// 		}, nil
// 	},
// },
// 		{
// 			name: "error in databse",
// 			args: args{
// 				ctx: context.Background(),
// 				jobData: models.Job{
// 					Cid:          1,
// 					Name:         "Software Engineer",
// 					Salary:       "20000",
// 					NoticePeriod: "3 months",
// 				},
// 				cid: 1,
// 			},
// 			want:    models.Job{},
// 			wantErr: true,
// 			mockResponse: func() (models.Job, error) {
// 				return models.Job{}, errors.New("could not creat job")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockResponse()).AnyTimes()

// 			s, err := NewService(mockRepo, &auth.Auth{})
// 			if err != nil {
// 				t.Errorf("error is initializing the repo layer")
// 				return
// 			}

// 			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewJobById(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		jid uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		want         models.Job
// 		wantErr      bool
// 		mockresponse func() (models.Job, error)
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "successful",
// 			args: args{
// 				ctx: context.Background(),
// 				jid: 6,
// 			},
// 			want: models.Job{
// 				Company: models.Company{
// 					Name: "Info",
// 				},
// 				Cid:          1,
// 				Name:         "Software Developer",
// 				Salary:       "20000",
// 				NoticePeriod: "3 months",
// 			},
// 			wantErr: false,
// 			mockresponse: func() (models.Job, error) {
// 				return models.Job{
// 					Company: models.Company{
// 						Name: "Info",
// 					},
// 					Cid:          1,
// 					Name:         "Software Developer",
// 					Salary:       "20000",
// 					NoticePeriod: "3 months",
// 				}, nil
// 			},
// 		},
// 		{
// 			name: "error in databse",
// 			args: args{
// 				ctx: context.Background(),
// 				jid: 15,
// 			},
// 			want:    models.Job{},
// 			wantErr: true,
// 			mockresponse: func() (models.Job, error) {
// 				return models.Job{}, errors.New("error test")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, tt.args.jid).Return(tt.mockresponse()).AnyTimes()
// 			s, err := NewService(mockRepo, &auth.Auth{})
// 			if err != nil {
// 				t.Errorf("error is initializing the repo layer")
// 				return
// 			}

// 			got, err := s.ViewJobById(tt.args.ctx, tt.args.jid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewJobById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewAllJobs(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		want         []models.Job
// 		wantErr      bool
// 		mockResponse func() ([]models.Job, error)
// 	}{

// 		{
// 			name: "error from databse",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			want:    nil,
// 			wantErr: true,
// 			mockResponse: func() ([]models.Job, error) {
// 				return nil, errors.New("could not find jobs in databse")
// 			},
// 		},
// 		// TODO: Add test cases.
// 		{
// 			name: "database success",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			want: []models.Job{
// 				{
// 					Cid:          1,
// 					Name:         "Software Engineer",
// 					Salary:       "20000",
// 					NoticePeriod: "3months",
// 				},
// 				{
// 					Cid:          1,
// 					Name:         "Softaware Trainee",
// 					Salary:       "20000",
// 					NoticePeriod: "3 months",
// 				},
// 			},
// 			wantErr: false,
// 			mockResponse: func() ([]models.Job, error) {
// 				return []models.Job{
// 					{
// 						Cid:          1,
// 						Name:         "Software Engineer",
// 						Salary:       "20000",
// 						NoticePeriod: "3 months",
// 					},
// 					{
// 						Cid:          1,
// 						Name:         "Software Trainee",
// 						Salary:       "20000",
// 						NoticePeriod: "3 months",
// 					},
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().FindAllJobs(tt.args.ctx).Return(tt.mockResponse()).AnyTimes()
// 			s, err := NewService(mockRepo, &auth.Auth{})
// 			if err != nil {
// 				t.Errorf("error in initializing service layer")
// 				return
// 			}
// 			got, err := s.ViewAllJobs(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewAllJobs() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewAllJobs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewJob(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		cid uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		want         []models.Job
// 		wantErr      bool
// 		mockResponse func() ([]models.Job, error)
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "success",
// 			args: args{
// 				ctx: context.Background(),
// 				cid: 1,
// 			},
// 			want: []models.Job{
// 				{
// 					Cid:          1,
// 					Name:         "software Engineer",
// 					Salary:       "20000",
// 					NoticePeriod: "3 months",
// 				},
// 				{
// 					Cid:          1,
// 					Name:         "Software Trainee",
// 					Salary:       "20000",
// 					NoticePeriod: "3 months",
// 				},
// 			},
// 			wantErr: false,
// 			mockResponse: func() ([]models.Job, error) {
// 				return []models.Job{
// 					{
// 						Cid:          1,
// 						Name:         "software Engineer",
// 						Salary:       "20000",
// 						NoticePeriod: "3 months",
// 					},
// 					{
// 						Cid:          1,
// 						Name:         "Software Trainee",
// 						Salary:       "20000",
// 						NoticePeriod: "3 months",
// 					},
// 				}, nil
// 			},
// 		},
// 		{
// 			name: "error from database",
// 			args: args{
// 				ctx: context.Background(),
// 				cid: 15,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 			mockResponse: func() ([]models.Job, error) {
// 				return nil, errors.New("could not find jobs")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().FindJob(tt.args.ctx, tt.args.cid).Return(tt.mockResponse()).AnyTimes()

// 			s, err := NewService(mockRepo, &auth.Auth{})
// 			if err != nil {
// 				t.Errorf("could not initialize service layer")
// 				return
// 			}
// 			got, err := s.ViewJob(tt.args.ctx, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewJob() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewJob() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
