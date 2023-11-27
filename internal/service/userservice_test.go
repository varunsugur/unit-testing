package service

import (
	"context"
	"golang/internal/auth"
	"golang/internal/cache"
	"golang/internal/models"
	"golang/internal/repository"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_VerifyUser(t *testing.T) {
	type args struct {
		ctx context.Context
		vu  models.VerifyUser
	}
	tests := []struct {
		name              string
		args              args
		wantErr           bool
		want              string
		mockRepoResponse  func() (models.User, error)
		mockCacheResponse func() error
	}{
		// TODO: Add test cases.
		// {
		// 	name: "error from repository layer",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		vu: models.VerifyUser{
		// 			UserEmail: "varun.sugur@gmail.com",
		// 			DOB:       "03/03/2000",
		// 		},
		// 	},
		// 	want:    "email is not found",
		// 	wantErr: true,
		// 	mockRepoResponse: func() (models.User, error) {
		// 		return models.User{}, errors.New("email is not found")
		// 	},
		// 	mockCacheResponse: func() error {
		// 		return nil
		// 	},
		// },
		// {
		// 	name: "error in date of birth",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		vu: models.VerifyUser{
		// 			UserEmail: "varun.sugur@gmail.com",
		// 			DOB:       "03/03/2000",
		// 		},
		// 	},
		// 	want:    "enter correct date of birth",
		// 	wantErr: true,
		// 	mockRepoResponse: func() (models.User, error) {
		// 		return models.User{
		// 			Name:  "varun",
		// 			Email: "varun.sugur@gmail.com",
		// 			DOB:   "02/02/2000",
		// 		}, nil
		// 	},
		// 	mockCacheResponse: func() error {
		// 		return nil
		// 	},
		// },
		{
			name: "error in sending mail",
			args: args{
				ctx: context.Background(),
				vu: models.VerifyUser{
					UserEmail: "varun.sugur@gmail.com",
					DOB:       "03/03/2000",
				},
			},
			want:    "could not send otp to verified email",
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Name:  "varun",
					Email: "varu.sugur@gmail.com",
					DOB:   "03/03/2000",
				}, nil
			},
			mockCacheResponse: func() error {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockCache := cache.NewMockCache(mc)

			mockCache.EXPECT().AddtoOTPCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
			mr := repository.NewMockUserRepo(mc)

			mr.EXPECT().VerifyUser(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			svc, err := NewService(mr, &auth.Auth{}, mockCache)
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}

			if err := svc.VerifyUser(tt.args.ctx, tt.args.vu); (err != nil) != tt.wantErr {
				t.Errorf("Service.VerifyUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func TestService_UpdatePassword(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		details models.ResetDetails
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		wantErr          bool
// 		want             string
// 		mockRepoResponse func() error
// 	}{
// 		// TODO: Add test cases.
// 		{

// 		}
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.s.UpdatePassword(tt.args.ctx, tt.args.details); (err != nil) != tt.wantErr {
// 				t.Errorf("Service.UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
