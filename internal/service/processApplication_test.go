package service

// import (
// 	"context"
// 	"golang/internal/models"
// 	"reflect"
// 	"testing"
// )

// func TestService_ProccessApplication(t *testing.T) {
// 	type args struct {
// 		ctx             context.Context
// 		applicationData []models.UserApplication
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		want         []models.UserApplication
// 		wantErr      bool
// 		mockResponse func()([]models.UserApplication, error)
// 	}{
// 		// TODO: Add test cases.

// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.s.ProccessApplication(tt.args.ctx, tt.args.applicationData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ProccessApplication() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ProccessApplication() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
