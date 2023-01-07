package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	e "github.com/natanaelb27/staycay-backend/internal/errors"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
	"github.com/natanaelb27/staycay-backend/internal/mocks"
	"github.com/natanaelb27/staycay-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAllHousesHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetAllHouses", MockQueryParams).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetAllHouses", MockQueryParams).Return([]*dtos.HouseDetailsWithPhotoDTO{}, &helpers.Pagination{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockQueryParams, h.GetAllHousesHandler)

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

func TestHandler_GetHouseByIDHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - House Not Found",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHouseByID", MockParamsHouseID, MockQueryParams).Return(nil, e.ErrHouseNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHouseByID", MockParamsHouseID, MockQueryParams).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHouseByID", MockParamsHouseID, MockQueryParams).Return(&dtos.HouseDetailsWithPhotoDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockParams, MiddlewareMockQueryParams, h.GetHouseByIDHandler)

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

func TestHandler_GetHostHousesHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - User Not Found",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHostHouses", MockUser.UserID, MockQueryParams).Return(nil, nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - User is not a Host",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHostHouses", MockUser.UserID, MockQueryParams).Return(nil, nil, e.ErrInvalidRoleNotHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHostHouses", MockUser.UserID, MockQueryParams).Return(nil, nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetHostHouses", MockUser.UserID, MockQueryParams).Return([]*dtos.HouseDetailsWithPhotoDTO{}, &helpers.Pagination{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockQueryParams, h.GetHostHousesHandler)

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

func TestHandler_GetAllHouseStatusesHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/status",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetAllHouseStatuses").Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/house/status",
			},
			mock: func(m *mocks.HouseService) {
				m.On("GetAllHouseStatuses").Return([]*dtos.HouseStatusDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.GET(tt.args.endPoint, h.GetAllHouseStatusesHandler)

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

func TestHandler_AddNewHouseHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		request      *dtos.AddNewHouseRequestDTO
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - Bad Request",
			houseService: mocks.NewHouseService(t),
			request:      &dtos.AddNewHouseRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:         "Failed - User Not Found",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddNewHouseRequestDTO{
				Name:        "House Name",
				Price:       50000,
				Description: "House Description",
				CityID:      1,
				MaxGuest:    5,
				Photo:       []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddNewHouse", MockUser.UserID, &dtos.AddNewHouseRequestDTO{
					Name:        "House Name",
					Price:       50000,
					Description: "House Description",
					CityID:      1,
					MaxGuest:    5,
					Photo:       []string{"image1"},
				}).Return(nil, []string{}, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - User is not a Host",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddNewHouseRequestDTO{
				Name:        "House Name",
				Price:       50000,
				Description: "House Description",
				CityID:      1,
				MaxGuest:    5,
				Photo:       []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddNewHouse", MockUser.UserID, &dtos.AddNewHouseRequestDTO{
					Name:        "House Name",
					Price:       50000,
					Description: "House Description",
					CityID:      1,
					MaxGuest:    5,
					Photo:       []string{"image1"},
				}).Return(nil, []string{}, e.ErrInvalidRoleNotHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddNewHouseRequestDTO{
				Name:        "House Name",
				Price:       50000,
				Description: "House Description",
				CityID:      1,
				MaxGuest:    5,
				Photo:       []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddNewHouse", MockUser.UserID, &dtos.AddNewHouseRequestDTO{
					Name:        "House Name",
					Price:       50000,
					Description: "House Description",
					CityID:      1,
					MaxGuest:    5,
					Photo:       []string{"image1"},
				}).Return(nil, []string{}, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddNewHouseRequestDTO{
				Name:        "House Name",
				Price:       50000,
				Description: "House Description",
				CityID:      1,
				MaxGuest:    5,
				Photo:       []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddNewHouse", MockUser.UserID, &dtos.AddNewHouseRequestDTO{
					Name:        "House Name",
					Price:       50000,
					Description: "House Description",
					CityID:      1,
					MaxGuest:    5,
					Photo:       []string{"image1"},
				}).Return(&models.House{}, []string{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, h.AddNewHouseHandler)

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

func TestHandler_UpdateHouseHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		request      interface{}
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - Bad Request",
			houseService: mocks.NewHouseService(t),
			request:      "invalid request",
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:         "Failed - User Not Found",
			houseService: mocks.NewHouseService(t),
			request: &dtos.UpdateHouseRequestDTO{
				Name: "House Name Updated",
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("UpdateHouse", MockParamsHouseID, MockUser.UserID, &dtos.UpdateHouseRequestDTO{
					Name: "House Name Updated",
				}).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - House Not Found",
			houseService: mocks.NewHouseService(t),
			request: &dtos.UpdateHouseRequestDTO{
				Name: "House Name Updated",
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("UpdateHouse", MockParamsHouseID, MockUser.UserID, &dtos.UpdateHouseRequestDTO{
					Name: "House Name Updated",
				}).Return(nil, e.ErrHouseNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - User is not a Host",
			houseService: mocks.NewHouseService(t),
			request: &dtos.UpdateHouseRequestDTO{
				Name: "House Name Updated",
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("UpdateHouse", MockParamsHouseID, MockUser.UserID, &dtos.UpdateHouseRequestDTO{
					Name: "House Name Updated",
				}).Return(nil, e.ErrInvalidRoleNotHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - User is not the Host",
			houseService: mocks.NewHouseService(t),
			request: &dtos.UpdateHouseRequestDTO{
				Name: "House Name Updated",
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("UpdateHouse", MockParamsHouseID, MockUser.UserID, &dtos.UpdateHouseRequestDTO{
					Name: "House Name Updated",
				}).Return(nil, e.ErrInvalidHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			request: &dtos.UpdateHouseRequestDTO{
				Name: "House Name Updated",
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("UpdateHouse", MockParamsHouseID, MockUser.UserID, &dtos.UpdateHouseRequestDTO{
					Name: "House Name Updated",
				}).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			request: &dtos.UpdateHouseRequestDTO{
				Name: "House Name Updated",
			},
			args: args{
				method:   http.MethodPatch,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("UpdateHouse", MockParamsHouseID, MockUser.UserID, &dtos.UpdateHouseRequestDTO{
					Name: "House Name Updated",
				}).Return(&models.House{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.PATCH(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.UpdateHouseHandler)

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

func TestHandler_DeleteHouseHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - User not Found",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - User is not a Host",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, e.ErrInvalidRoleNotHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - User is not the Host",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, e.ErrInvalidHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - Cannot Delete due to Upcoming Res",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, e.ErrDeleteUpcomingReservationError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:         "Failed - House not Found",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, e.ErrHouseNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/:houseId",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(&models.House{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.DELETE(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.DeleteHouseHandler)

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

func TestHandler_AddHousePhotoHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		request      interface{}
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - Bad Request",
			houseService: mocks.NewHouseService(t),
			request:      "invalid request",
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock:  func(m *mocks.HouseService) {},
			wants: http.StatusBadRequest,
		},
		{
			name:         "Failed - User not Found",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddHousePhotoRequestDTO{
				Photo: []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddHousePhoto", MockParamsHouseID, MockUser.UserID, &dtos.AddHousePhotoRequestDTO{
					Photo: []string{"image1"},
				}).Return([]string{}, e.ErrUserNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - User is not a Host",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddHousePhotoRequestDTO{
				Photo: []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddHousePhoto", MockParamsHouseID, MockUser.UserID, &dtos.AddHousePhotoRequestDTO{
					Photo: []string{"image1"},
				}).Return([]string{}, e.ErrInvalidRoleNotHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - User is not the Host",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddHousePhotoRequestDTO{
				Photo: []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddHousePhoto", MockParamsHouseID, MockUser.UserID, &dtos.AddHousePhotoRequestDTO{
					Photo: []string{"image1"},
				}).Return([]string{}, e.ErrInvalidHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddHousePhotoRequestDTO{
				Photo: []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddHousePhoto", MockParamsHouseID, MockUser.UserID, &dtos.AddHousePhotoRequestDTO{
					Photo: []string{"image1"},
				}).Return([]string{}, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			request: &dtos.AddHousePhotoRequestDTO{
				Photo: []string{"image1"},
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/host/house",
			},
			mock: func(m *mocks.HouseService) {
				m.On("AddHousePhoto", MockParamsHouseID, MockUser.UserID, &dtos.AddHousePhotoRequestDTO{
					Photo: []string{"image1"},
				}).Return([]string{"image1"}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.AddHousePhotoHandler)

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

func TestHandler_DeleteHousePhotoHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name         string
		houseService *mocks.HouseService
		args         args
		mock         func(*mocks.HouseService)
		wants        int
	}{
		{
			name:         "Failed - House Photo not Found",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/photo",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHousePhoto", MockParamsHousePhotoID, MockUser.UserID).Return(nil, e.ErrHousePhotoNotFoundError)
			},
			wants: http.StatusNotFound,
		},
		{
			name:         "Failed - House must have at least one Photo",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/photo",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHousePhoto", MockParamsHousePhotoID, MockUser.UserID).Return(nil, e.ErrHousePhotoMustMoreThanOneError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:         "Failed - User is not the Host",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/photo",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHousePhoto", MockParamsHousePhotoID, MockUser.UserID).Return(nil, e.ErrInvalidHostError)
			},
			wants: http.StatusForbidden,
		},
		{
			name:         "Failed - Internal Server Error",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/photo",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHousePhoto", MockParamsHousePhotoID, MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:         "Success",
			houseService: mocks.NewHouseService(t),
			args: args{
				method:   http.MethodDelete,
				endPoint: "/host/house/photo",
			},
			mock: func(m *mocks.HouseService) {
				m.On("DeleteHousePhoto", MockParamsHousePhotoID, MockUser.UserID).Return(&models.HousePhoto{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				HouseService: tt.houseService,
			})

			tt.mock(tt.houseService)

			r := gin.Default()

			r.DELETE(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.DeleteHousePhotoHandler)

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

// func TestHandler_DeleteHouseHandler(t *testing.T) {
// 	type args struct {
// 		method   string
// 		endPoint string
// 	}
// 	tests := []struct {
// 		name         string
// 		houseService *mocks.HouseService
// 		args         args
// 		mock         func(*mocks.HouseService)
// 		wants        int
// 	}{
// 		{
// 			name:         "Failed - Internal Server Error",
// 			houseService: mocks.NewHouseService(t),
// 			args: args{
// 				method:   http.MethodDelete,
// 				endPoint: "/host/house",
// 			},
// 			mock: func(m *mocks.HouseService) {
// 				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(nil, errors.New("unknown error"))
// 			},
// 			wants: http.StatusInternalServerError,
// 		},
// 		{
// 			name:         "Success",
// 			houseService: mocks.NewHouseService(t),
// 			args: args{
// 				method:   http.MethodDelete,
// 				endPoint: "/host/house",
// 			},
// 			mock: func(m *mocks.HouseService) {
// 				m.On("DeleteHouse", MockParamsHouseID, MockUser.UserID).Return(&models.House{}, nil)
// 			},
// 			wants: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := New(&HandlerConfig{
// 				HouseService: tt.houseService,
// 			})

// 			tt.mock(tt.houseService)

// 			r := gin.Default()

// 			r.DELETE(tt.args.endPoint, MiddlewareMockUser, MiddlewareMockParams, h.DeleteHouseHandler)

// 			req, _ := http.NewRequest(
// 				tt.args.method,
// 				tt.args.endPoint,
// 				nil,
// 			)

// 			w := httptest.NewRecorder()

// 			r.ServeHTTP(w, req)

// 			assert.Equal(t, tt.wants, w.Code)
// 		})
// 	}
// }
