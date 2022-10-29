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

func TestHandler_AddReservationHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name               string
		reservationService *mocks.ReservationService
		request            *dtos.AddReservationRequestDTO
		args               args
		mock               func(*mocks.ReservationService)
		wants              int
	}{
		{
			name:               "Failed - Bad Request",
			reservationService: mocks.NewReservationService(t),
			request:            &dtos.AddReservationRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - User not Found",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:               "Failed - House not Found",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrHouseNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:               "Failed - User is the Host",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrInvalidReservationUserIsHostError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - Check In Date is Later than the Check Out Date",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrInvalidReservationDateLaterError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - Check In Date is the same as the Check Out Date",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrInvalidReservationDateSameError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - Check In Date is Earlier than Today",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrInvalidReservationBeforeNowError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - Conflicting Reservation Date",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrNotValidReservationDateConflict)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - Insufficient Balance",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, e.ErrInsufficientBalanceError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:               "Failed - Internal Server Error",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:               "Success",
			reservationService: mocks.NewReservationService(t),
			request: &dtos.AddReservationRequestDTO{
				HouseID:      1,
				CheckInDate:  "2022-11-11",
				CheckOutDate: "2022-11-12",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("AddReservation", MockUser.UserID, &dtos.AddReservationRequestDTO{
					HouseID:      1,
					CheckInDate:  "2022-11-11",
					CheckOutDate: "2022-11-12",
				}).Return(&models.Reservation{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				ReservationService: tt.reservationService,
			})

			tt.mock(tt.reservationService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, h.AddReservationHandler)

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

func TestHandler_GetHouseReservationsByUserHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name               string
		reservationService *mocks.ReservationService
		args               args
		mock               func(*mocks.ReservationService)
		wants              int
	}{
		{
			name:               "Failed - Internal Server Error",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByUser", MockUser.UserID, MockQueryParams).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:               "Success",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/user/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByUser", MockUser.UserID, MockQueryParams).Return([]*dtos.HouseReservationsDTO{}, &helpers.Pagination{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				ReservationService: tt.reservationService,
			})

			tt.mock(tt.reservationService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockQueryParams, h.GetHouseReservationsByUserHandler)

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

func TestHandler_GetHouseReservationsByHostHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name               string
		reservationService *mocks.ReservationService
		args               args
		mock               func(*mocks.ReservationService)
		wants              int
	}{
		{
			name:               "Failed - Internal Server Error",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/host/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByHost", MockUser.UserID, MockQueryParams).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:               "Success",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/host/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByHost", MockUser.UserID, MockQueryParams).Return([]*dtos.HouseReservationsDTO{}, &helpers.Pagination{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				ReservationService: tt.reservationService,
			})

			tt.mock(tt.reservationService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockQueryParams, h.GetHouseReservationsByHostHandler)

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

func TestHandler_GetHouseReservationsByHouseHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name               string
		reservationService *mocks.ReservationService
		args               args
		mock               func(*mocks.ReservationService)
		wants              int
	}{
		{
			name:               "Failed - Internal Server Error",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByHouse", MockParamsHouseID, MockQueryParams).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:               "Success",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/reservation",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByHouse", MockParamsHouseID, MockQueryParams).Return([]*dtos.HouseReservationsDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				ReservationService: tt.reservationService,
			})

			tt.mock(tt.reservationService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockParams, MiddlewareMockQueryParams, h.GetHouseReservationsByHouseHandler)

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

func TestHandler_GetHouseReservationsByIDHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name               string
		reservationService *mocks.ReservationService
		args               args
		mock               func(*mocks.ReservationService)
		wants              int
	}{
		{
			name:               "Failed - Internal Server Error",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/reservation/:reservationId",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByID", MockParamsReservationID, MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:               "Failed - Reservation not Found",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/reservation/:reservationId",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByID", MockParamsReservationID, MockUser.UserID).Return(nil, e.ErrReservationNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:               "Failed - User not Found",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/reservation/:reservationId",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByID", MockParamsReservationID, MockUser.UserID).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:               "Failed - House not Found",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/reservation/:reservationId",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByID", MockParamsReservationID, MockUser.UserID).Return(nil, e.ErrHouseNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:               "Failed - User is Neither the Host or the Customer",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/reservation/:reservationId",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByID", MockParamsReservationID, MockUser.UserID).Return(nil, e.ErrInvalidHostAndUserReservationError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:               "Success",
			reservationService: mocks.NewReservationService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/reservation/:reservationId",
			},
			mock: func(m *mocks.ReservationService) {
				m.On("GetHouseReservationsByID", MockParamsReservationID, MockUser.UserID).Return(&dtos.HouseReservationsDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				ReservationService: tt.reservationService,
			})

			tt.mock(tt.reservationService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.GetHouseReservationsByIDHandler)

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
