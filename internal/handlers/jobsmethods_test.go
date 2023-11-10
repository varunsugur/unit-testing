package handlers

import (
	"context"
	"strings"

	"golang/internal/auth"
	"golang/internal/middlewares"

	"golang/internal/service"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"

	// "github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func TestHandler_AddJobs(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		// TODO: Add test cases.

		// {
		// 	name: "missing trace id",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// 		c.Request = httpRequest

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"error":"Internal Server Error"}`,
		// },

		// {
		// 	name: "missing jwt token",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
		// 		httpRequest = httpRequest.WithContext(ctx)

		// 		c.Request = httpRequest

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusUnauthorized,
		// 	expectedResponse:   `{"error":"Unauthorized"}`,
		// },

		// {
		// 	name: "invalid job id",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest

		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockUserService(mc)

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `"Bad Request"`,
		// },
		{
			name: "error in validating the json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"name":"Software-Trainee",
					"budget":"5lack",
					"minNoticePeriod":"7",
					"maxNoticePeriod":"30",
					"jobLocation":[1,7 ],
					"technology":[1,5],
					"description": "Develop-Software Application and Code-Review",
					"minExp": "0",
					"maxExp": "4",
					"qualification":[3,4],
					"shift":[1],
					"jobType":"Full-Time
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := service.NewMockUserService(mc)

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `"Bad Request"`,
		},

		// {
		// 	name: "error ferom service layer",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`
		// 		{
		// 			"name":"Software-Trainee",
		// 			"budget":"5lack",
		// 			"minNoticePeriod":"7",
		// 			"maxNoticePeriod":"30",
		// 			"jobLocation":[1,7 ],
		// 			"technology":[1,5],
		// 			"description": "Develop-Software Application and Code-Review",
		// 			"minExp": "0",
		// 			"maxExp": "4",
		// 			"qualification":[3,4],
		// 			"shift":[1],
		// 			"jobType":"Full-Time"
		// 		}`,
		// 		))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		c.Request = httpRequest

		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockUserService(mc)

		// 		ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.ResponseJob{}, errors.New("test error from mock function")).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `"Bad Request"`,
		// },

		// {
		// 	name: "success from service layer",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{
		// 			"name":"Software-Trainee",
		// 			"budget":"5lack",
		// 			"minNoticePeriod":"7",
		// 			"maxNoticePeriod":"30",
		// 			"jobLocation":[1,7 ],
		// 			"technology":[1,5],
		// 			"description": "Develop-Software Application and Code-Review",
		// 			"minExp": "0",
		// 			"maxExp": "4",
		// 			"qualification":[3,4],
		// 			"shift":[1],
		// 			"jobType":"Full-Time"
		// 		}`,
		// 		))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewares.TraceIdKey, "123")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

		// 		c.Request = httpRequest

		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockUserService(mc)
		// 		ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.ResponseJob{
		// 			Id: 1,
		// 		}, nil).AnyTimes()
		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `{"responseId":1}`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := Handler{
				service: ms,
			}
			h.AddJobs(c)
			// assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
