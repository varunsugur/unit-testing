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

	"github.com/redis/go-redis/v9"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_ProccessApplication(t *testing.T) {
	type args struct {
		ctx             context.Context
		applicationData []models.UserApplication
	}
	tests := []struct {
		name              string
		args              args
		want              []models.UserApplication
		wantErr           bool
		mockRepoResponse  func() (models.Job, error)
		mockcacheResponse func() (models.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "not found in redis and in databse",
			args: args{
				ctx: context.Background(),
				applicationData: []models.UserApplication{
					{
						Name:    "varun",
						College: "nie",
						Jid:     1,
						Job: models.JobRequest{
							Jobname:        "Spoftware Developer",
							NoticePeriod:   "7",
							Experience:     "2",
							Location:       []uint{1, 2},
							Technology:     []uint{1, 2},
							Qualifications: []uint{2, 3},
							Shift:          []uint{1},
							JobType:        "Full-Time",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error from cache mock")
			},
			mockcacheResponse: func() (models.Job, error) {
				return models.Job{}, redis.Nil
			},
		},
		{
			name: "not found in redis and in database",
			args: args{
				ctx: context.Background(),
				applicationData: []models.UserApplication{
					{
						Name:    "varun",
						College: "nie",
						Jid:     1,
						Job: models.JobRequest{
							Jobname:        "Software Developer",
							NoticePeriod:   "7",
							Experience:     "2",
							Location:       []uint{1, 2},
							Technology:     []uint{1, 2},
							Qualifications: []uint{2, 3},
							Shift:          []uint{1},
							JobType:        "Full-Time",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Model: gorm.Model{
						ID: 1,
					},
					Name:            "Software Developer",
					Budget:          "500000",
					MinNoticePeriod: "7",
					MaxNoticePeriod: "30",
					JobLocation: []models.Location{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Technology: []models.Technology{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					Description: "you are going to develop a app ",
					MinExp:      "2",
					MaxExp:      "5",
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					JobType: "Full-Time",
				}, nil
			},
			mockcacheResponse: func() (models.Job, error) {
				return models.Job{}, redis.Nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockCash := cache.NewMockCache(mc)
			mockCash.EXPECT().GetTheCacheData(gomock.Any(), gomock.Any()).Return(tt.mockcacheResponse()).AnyTimes()
			mockCash.EXPECT().AddToCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test error")).AnyTimes()

			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().GetTheJobData(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{}, mockCash)
			if err != nil {
				t.Error("could not initialize the mock service layer")
				return
			}

			got, err := s.ProccessApplication(tt.args.ctx, tt.args.applicationData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ProccessApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ProccessApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
