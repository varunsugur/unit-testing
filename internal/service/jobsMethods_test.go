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
