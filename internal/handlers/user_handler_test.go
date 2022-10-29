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

func TestHandler_GetUserProfileHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		userService *mocks.UserService
		args        args
		mock        func(*mocks.UserService)
		wants       int
	}{
		{
			name:        "Failed - User not Found",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("GetUserProfile", MockUser.UserID).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:        "Failed - Internal Server Error",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("GetUserProfile", MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("GetUserProfile", MockUser.UserID).Return(&dtos.UserDetailsDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				UserService: tt.userService,
			})

			tt.mock(tt.userService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, h.GetUserProfileHandler)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				nil,
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}

func TestHandler_GetUserProfileByUserIDHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		userService *mocks.UserService
		args        args
		mock        func(*mocks.UserService)
		wants       int
	}{
		{
			name:        "Failed - User not Found",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("GetUserProfileByUserID", MockParamsUserID).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:        "Failed - Internal Server Error",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("GetUserProfileByUserID", MockParamsUserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("GetUserProfileByUserID", MockParamsUserID).Return(&dtos.UserDetailsDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				UserService: tt.userService,
			})

			tt.mock(tt.userService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockParams, h.GetUserProfileByUserIDHandler)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				nil,
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}

func TestHandler_UpdateUserRoleHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		userService *mocks.UserService
		args        args
		mock        func(*mocks.UserService)
		wants       int
	}{
		{
			name:        "Failed - Internal Server Error",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user/role",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserRoleToHost", MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			userService: mocks.NewUserService(t),
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user/role",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserRoleToHost", MockUser.UserID).Return(&models.User{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				UserService: tt.userService,
			})

			tt.mock(tt.userService)

			r := gin.Default()

			r.PATCH(tt.args.endPoint, MiddlewareMockUser, h.UpdateUserRoleHandler)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				nil,
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}

func TestHandler_UpdateUserProfileHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		userService *mocks.UserService
		request     interface{}
		args        args
		mock        func(*mocks.UserService)
		wants       int
	}{
		{
			name:        "Failed - Bad Request",
			userService: mocks.NewUserService(t),
			request:     "invalid request",
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - User not Found",
			userService: mocks.NewUserService(t),
			request: &dtos.UpdateUserProfileRequestDTO{
				Name:        "Name Updated",
				OldPassword: "OldPassword123",
				Password:    "NewPassword123",
				Address:     "Street Updated",
				CityID:      2,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserProfile", MockUser.UserID, &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "NewPassword123",
					Address:     "Street Updated",
					CityID:      2,
				}).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:        "Failed - Invalid Password Input",
			userService: mocks.NewUserService(t),
			request: &dtos.UpdateUserProfileRequestDTO{
				Name:        "Name Updated",
				OldPassword: "OldPassword123",
				Password:    "NewPassword",
				Address:     "Street Updated",
				CityID:      2,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserProfile", MockUser.UserID, &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "NewPassword",
					Address:     "Street Updated",
					CityID:      2,
				}).Return(nil, e.ErrInvalidPasswordInputError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Old Password does not Match",
			userService: mocks.NewUserService(t),
			request: &dtos.UpdateUserProfileRequestDTO{
				Name:        "Name Updated",
				OldPassword: "OldPassword123",
				Password:    "NewPassword123",
				Address:     "Street Updated",
				CityID:      2,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserProfile", MockUser.UserID, &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "NewPassword123",
					Address:     "Street Updated",
					CityID:      2,
				}).Return(nil, e.ErrOldPasswordNotMatchError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Internal Server Error",
			userService: mocks.NewUserService(t),
			request: &dtos.UpdateUserProfileRequestDTO{
				Name:        "Name Updated",
				OldPassword: "OldPassword123",
				Password:    "NewPassword123",
				Address:     "Street Updated",
				CityID:      2,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserProfile", MockUser.UserID, &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "NewPassword123",
					Address:     "Street Updated",
					CityID:      2,
				}).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			userService: mocks.NewUserService(t),
			request: &dtos.UpdateUserProfileRequestDTO{
				Name:        "Name Updated",
				OldPassword: "OldPassword123",
				Password:    "NewPassword123",
				Address:     "Street Updated",
				CityID:      2,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/user",
			},
			mock: func(m *mocks.UserService) {
				m.On("UpdateUserProfile", MockUser.UserID, &dtos.UpdateUserProfileRequestDTO{
					Name:        "Name Updated",
					OldPassword: "OldPassword123",
					Password:    "NewPassword123",
					Address:     "Street Updated",
					CityID:      2,
				}).Return(&models.User{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				UserService: tt.userService,
			})

			tt.mock(tt.userService)

			r := gin.Default()

			r.PATCH(tt.args.endPoint, MiddlewareMockUser, h.UpdateUserProfileHandler)

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
