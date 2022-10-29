package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAllCitiesHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		cityService *mocks.CityService
		args        args
		mock        func(*mocks.CityService)
		wants       int
	}{
		{
			name:        "Failed - Internal Server Error",
			cityService: mocks.NewCityService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/city",
			},
			mock: func(m *mocks.CityService) {
				m.On("GetAllCities").Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			cityService: mocks.NewCityService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/city",
			},
			mock: func(m *mocks.CityService) {
				m.On("GetAllCities").Return([]*dtos.CitiesDTO{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				CityService: tt.cityService,
			})

			tt.mock(tt.cityService)

			r := gin.Default()

			r.GET(tt.args.endPoint, h.GetAllCitiesHandler)

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
