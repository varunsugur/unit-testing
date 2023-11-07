package service

import (
	"context"
	"golang/internal/models"
	"reflect"
	"testing"
)

func TestService_UserSignup(t *testing.T) {
	type args struct {
		ctx context.Context
		nu  models.NewUser
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UserSignup(tt.args.ctx, tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UserSignin(t *testing.T) {
	type args struct {
		ctx  context.Context
		user models.UserLogin
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UserSignin(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.UserSignin() = %v, want %v", got, tt.want)
			}
		})
	}
}
