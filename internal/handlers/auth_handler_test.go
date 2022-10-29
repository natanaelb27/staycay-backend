package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_RegisterHandler(t *testing.T) {
	token := "jwttoken"

	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		authService *mocks.AuthService
		request     *dtos.RegisterRequestDTO
		args        args
		mock        func(*mocks.AuthService)
		wants       int
	}{
		{
			name:        "Failed - Bad Request",
			authService: mocks.NewAuthService(t),
			request:     &dtos.RegisterRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/register",
			},
			mock: func(m *mocks.AuthService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Invalid Email Input",
			authService: mocks.NewAuthService(t),
			request: &dtos.RegisterRequestDTO{
				Name:     "Name",
				Email:    "email123@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				IsAdmin:  false,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/register",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Register", &dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "email123@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				}).Return(nil, nil, e.ErrInvalidEmailInputError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Invalid Password Input",
			authService: mocks.NewAuthService(t),
			request: &dtos.RegisterRequestDTO{
				Name:     "Name",
				Email:    "email123@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				IsAdmin:  false,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/register",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Register", &dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "email123@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				}).Return(nil, nil, e.ErrInvalidPasswordInputError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Duplicate Email",
			authService: mocks.NewAuthService(t),
			request: &dtos.RegisterRequestDTO{
				Name:     "Name",
				Email:    "email123@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				IsAdmin:  false,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/register",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Register", &dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "email123@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				}).Return(nil, nil, e.ErrDuplicateEmailError)
			},
			wants: http.StatusConflict,
		},
		{
			name:        "Failed - Internal Server Error",
			authService: mocks.NewAuthService(t),
			request: &dtos.RegisterRequestDTO{
				Name:     "Name",
				Email:    "email123@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				IsAdmin:  false,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/register",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Register", &dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "email123@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				}).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			authService: mocks.NewAuthService(t),
			request: &dtos.RegisterRequestDTO{
				Name:     "Name",
				Email:    "email123@email.com",
				Password: "Password123",
				Address:  "Badak Street",
				CityID:   1,
				IsAdmin:  false,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/register",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Register", &dtos.RegisterRequestDTO{
					Name:     "Name",
					Email:    "email123@email.com",
					Password: "Password123",
					Address:  "Badak Street",
					CityID:   1,
					IsAdmin:  false,
				}).Return(&models.User{}, &token, nil)
			},
			wants: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				AuthService: tt.authService,
			})

			tt.mock(tt.authService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, h.RegisterHandler)

			jsonValue, _ := json.Marshal(tt.request)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				bytes.NewBuffer(jsonValue),
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}

func TestHandler_LoginHandler(t *testing.T) {
	token := "jwttoken"

	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		authService *mocks.AuthService
		request     *dtos.LoginRequestDTO
		args        args
		mock        func(*mocks.AuthService)
		wants       int
	}{
		{
			name:        "Failed - Bad Request",
			authService: mocks.NewAuthService(t),
			request:     &dtos.LoginRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/login",
			},
			mock: func(m *mocks.AuthService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Credential Error",
			authService: mocks.NewAuthService(t),
			request: &dtos.LoginRequestDTO{
				Email:    "email123@email.com",
				Password: "Password123",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/login",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Login", &dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				}).Return(nil, nil, e.ErrCredentialError)
			},
			wants: http.StatusUnauthorized,
		},
		{
			name:        "Failed - Internal Server Error",
			authService: mocks.NewAuthService(t),
			request: &dtos.LoginRequestDTO{
				Email:    "email123@email.com",
				Password: "Password123",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/login",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Login", &dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				}).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			authService: mocks.NewAuthService(t),
			request: &dtos.LoginRequestDTO{
				Email:    "email123@email.com",
				Password: "Password123",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/login",
			},
			mock: func(m *mocks.AuthService) {
				m.On("Login", &dtos.LoginRequestDTO{
					Email:    "email123@email.com",
					Password: "Password123",
				}).Return(&models.User{}, &token, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				AuthService: tt.authService,
			})

			tt.mock(tt.authService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, h.LoginHandler)

			jsonValue, _ := json.Marshal(tt.request)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				bytes.NewBuffer(jsonValue),
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}
