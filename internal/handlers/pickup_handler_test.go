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
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_RequestPickupHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name          string
		pickupService *mocks.PickupService
		request       *dtos.RequestPickupRequestDTO
		args          args
		mock          func(*mocks.PickupService)
		wants         int
	}{
		{
			name:          "Failed - Bad Request",
			pickupService: mocks.NewPickupService(t),
			request:       &dtos.RequestPickupRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - Reservation not Found",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrReservationNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:          "Failed - User not Found",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:          "Failed - Reservation already Registered for Pickup",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrDuplicatePickupReservationError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - User is not the Customer",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrInvalidUserReservationError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:          "Failed - Pickup is not Available for Completed Reservation",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrPickupCompletedReservation)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - Pickup is not Available for Ongoing Reservation",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrPickupOngoingReservation)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - Insufficient Balance",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, e.ErrInsufficientBalanceError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - Internal Server Error",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:          "Success",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.RequestPickupRequestDTO{
				PickupAddress: "Main Street",
				CityID:        1,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("RequestPickup", MockParamsReservationID, MockUser.UserID, &dtos.RequestPickupRequestDTO{
					PickupAddress: "Main Street",
					CityID:        1,
				}).Return(&models.Pickup{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				PickupService: tt.pickupService,
			})

			tt.mock(tt.pickupService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.RequestPickupHandler)

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

func TestHandler_GetPickupDetailsByReservationHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name          string
		pickupService *mocks.PickupService
		args          args
		mock          func(*mocks.PickupService)
		wants         int
	}{
		{
			name:          "Failed - Reservation not Found",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetPickupDetailsByReservation", MockParamsReservationID, MockUser.UserID).Return(nil, e.ErrReservationNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:          "Failed - Internal Server Error",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetPickupDetailsByReservation", MockParamsReservationID, MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:          "Success",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetPickupDetailsByReservation", MockParamsReservationID, MockUser.UserID).Return(&dtos.GetPickupDetailsByReservationResponseDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				PickupService: tt.pickupService,
			})

			tt.mock(tt.pickupService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.GetPickupDetailsByReservationHandler)

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

func TestHandler_AdminUpdatePickupStatusHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name          string
		pickupService *mocks.PickupService
		request       interface{}
		args          args
		mock          func(*mocks.PickupService)
		wants         int
	}{
		{
			name:          "Failed - Bad Request",
			pickupService: mocks.NewPickupService(t),
			request:       "invalid request",
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - User not Found",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:          "Failed - Reservation not Found",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(nil, e.ErrReservationNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:          "Failed - Cannot Change Pickup Status: Check In Date is not Today",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(nil, e.ErrInvalidCheckInDateUpdatePickupStatusError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - User Role is not Admin",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(nil, e.ErrInvalidRoleNotAdminError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:          "Failed - Pickup not Found",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(nil, e.ErrPickupNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:          "Failed - Internal Server Error",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:          "Success",
			pickupService: mocks.NewPickupService(t),
			request: &dtos.AdminUpdatePickupStatusRequestDTO{
				PickupStatusID: 3,
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/admin/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("AdminUpdatePickupStatus", 3, MockParamsPickupID, MockUser.UserID).Return(&models.Pickup{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				PickupService: tt.pickupService,
			})

			tt.mock(tt.pickupService)

			r := gin.Default()

			r.PATCH(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.AdminUpdatePickupStatusHandler)

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

func TestHandler_GetAllPickupStatusesHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name          string
		pickupService *mocks.PickupService
		args          args
		mock          func(*mocks.PickupService)
		wants         int
	}{
		{
			name:          "Failed - Internal Server Error",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup/status",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetAllPickupStatuses").Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:          "Success",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup/status",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetAllPickupStatuses").Return([]*dtos.PickupStatusesDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				PickupService: tt.pickupService,
			})

			tt.mock(tt.pickupService)

			r := gin.Default()

			r.GET(tt.args.endPoint, h.GetAllPickupStatusesHandler)

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

func TestHandler_GetAllPickupDetailsHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name          string
		pickupService *mocks.PickupService
		args          args
		mock          func(*mocks.PickupService)
		wants         int
	}{
		{
			name:          "Failed - Internal Server Error",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetAllPickupDetails", MockQueryParams).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:          "Success",
			pickupService: mocks.NewPickupService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/pickup",
			},
			mock: func(m *mocks.PickupService) {
				m.On("GetAllPickupDetails", MockQueryParams).Return([]*dtos.PickupDetailsDTO{}, &helpers.Pagination{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				PickupService: tt.pickupService,
			})

			tt.mock(tt.pickupService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockQueryParams, h.GetAllPickupDetailsHandler)

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
