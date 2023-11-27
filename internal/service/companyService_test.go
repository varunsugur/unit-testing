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

func TestService_AddCompanyDetails(t *testing.T) {
	type args struct {
		ctx  context.Context
		data models.Company
	}
	tests := []struct {
		name         string
		args         args
		want         models.Company
		wantErr      bool
		mockResponse func() (models.Company, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				data: models.Company{
					Name:     "Info",
					Location: "Banagalore",
					Field:    "IT",
				},
			},
			want: models.Company{
				Name:     "Info",
				Location: "Bangalore",
				Field:    "IT",
			},
			wantErr: false,
			mockResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Info",
					Location: "Bangalore",
					Field:    "IT",
				}, nil
			},
		},
		{
			name: "error in databse",
			args: args{
				ctx: context.Background(),
				data: models.Company{
					Name:     "Info",
					Location: "Bnagalore",
					Field:    "IT",
				},
			},
			want:    models.Company{},
			wantErr: true,
			mockResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("could not creat company")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockrepo := repository.NewMockUserRepo(mc)
			mockrepo.EXPECT().CreatCompany(tt.args.ctx, tt.args.data).Return(tt.mockResponse).AnyTimes()

			s, err := NewService(mockrepo, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("error in initializing repo layer")
				return
			}

			got, err := s.AddCompanyDetails(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllCompanies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		args         args
		want         []models.Company
		wantErr      bool
		mockResponse func() ([]models.Company, error)
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Company{
				{
					Name:     "Info",
					Location: "Bangalore",
					Field:    "IT",
				},
				{
					Name:     "Wipro",
					Location: "Bangalore",
					Field:    "It",
				},
			},
			wantErr: false,
			mockResponse: func() ([]models.Company, error) {
				return []models.Company{
					{
						Name:     "Info",
						Location: "Bangalore",
						Field:    "IT",
					},
					{
						Name:     "Wipro",
						Location: "Bangalore",
						Field:    "It",
					},
				}, nil
			},
		},
		{
			name: "error in database",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockResponse: func() ([]models.Company, error) {
				return nil, errors.New("test error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().ViewCompanies(tt.args.ctx).Return(tt.mockResponse).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("error in initializing service layer")
				return
			}

			got, err := s.ViewAllCompanies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewCompanyDetail(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name         string
		args         args
		want         models.Company
		wantErr      bool
		mockResponse func() (models.Company, error)
	}{
		// TODO: Add test cases.
		{
			name: "success from the database",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: models.Company{
				Name:     "Info",
				Location: "Bangalore",
				Field:    "IT",
			},
			wantErr: false,
			mockResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Info",
					Location: "Bangalore",
					Field:    "IT",
				}, nil
			},
		},
		{
			name: "error from the database",
			args: args{
				ctx: context.Background(),
				cid: 15,
			},
			want:    models.Company{},
			wantErr: true,
			mockResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("test error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().ViewCompanyById(tt.args.ctx, tt.args.cid).Return(tt.mockResponse).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{}, &cache.MockCache{})
			if err != nil {
				t.Errorf("error in initializing service layer")
				return
			}

			got, err := s.ViewCompanyDetail(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewCompanyDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewCompanyDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
