package handlers

import (
	"context"
	"errors"
	"strings"

	// "errors"
	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/models"

	// "golang/internal/models"
	// "context"

	"golang/internal/service"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func TestHandler_ProcessApplication(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		// TODO: Add test cases.
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "misssing jwt token",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)

				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error in validating json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`[
					{
						"name": "varun",
						"college": "nie",
						"jid": 10,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "5,
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "0",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "rahul",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								3
							],
							"experience": "2",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "sagar",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "7",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "3",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "jeevan",
						"college": "nie",
						"jid": 11,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "1",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "rohit",
						"college": "nie",
						"jid": 11,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "18",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "0",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "mohit",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "2",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					}
				]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)

				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := service.NewMockUserService(mc)

				return c, rr, ms

			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"please provide valid job details"}`,
		},
		{
			name: "error from service layer",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`[
					{
						"name": "varun",
						"college": "nie",
						"jid": 10,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "5",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "0",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "rahul",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								3
							],
							"experience": "2",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
				]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := service.NewMockUserService(mc)

				ms.EXPECT().ProccessApplication(c.Request.Context(), gomock.Any()).Return([]models.UserApplication{}, errors.New("test service error")).AnyTimes()

				return c, rr, ms

			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"please provide valid job details"}`,
		},
		{
			name: "success from service layer",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`[
					{
						"name": "varun",
						"college": "nie",
						"jid": 10,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "5",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "0",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "rahul",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								3
							],
							"experience": "2",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "sagar",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "7",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "3",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "jeevan",
						"college": "nie",
						"jid": 11,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "1",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "rohit",
						"college": "nie",
						"jid": 11,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "18",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "0",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					},
					{
						"name": "mohit",
						"college": "nie",
						"jid": 12,
						"jobRequest": {
							"jobName": "Software-Developer",
							"noticePeriod": "8",
							"location": [
								1,
								7
							],
							"technology": [
								1,
								5
							],
							"experience": "2",
							"qualifications": [
								3,
								4
							],
							"shifts": [
								1
							],
							"jobtype": "Full-Time"
						}
					}
				]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)

				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := service.NewMockUserService(mc)

				ms.EXPECT().ProccessApplication(c.Request.Context(), gomock.Any()).Return([]models.UserApplication{
					{
						Name:    "varun",
						College: "nie",
					},
					{
						Name:    "rahul",
						College: "nie",
					},
				}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"name":"varun","college":"nie","jid":0,"jobRequest":{"jobName":"","noticePeriod":"","location":null,"technology":null,"experience":"","qualifications":null,"shifts":null,"jobtype":""}},{"name":"rahul","college":"nie","jid":0,"jobRequest":{"jobName":"","noticePeriod":"","location":null,"technology":null,"experience":"","qualifications":null,"shifts":null,"jobtype":""}}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := Handler{
				service: ms,
			}
			h.ProcessApplication(c)

			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
