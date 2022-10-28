package services

import (
	"errors"
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_cityService_NewCityService(t *testing.T) {
	cityRepository := mocks.NewCityRepository(t)
	cityService := NewCityService(
		&CSConfig{
			CityRepository: cityRepository,
		},
	)

	assert.NotNil(t, cityService)
}

func Test_cityService_GetAllCities(t *testing.T) {
	tests := []struct {
		name           string
		cityRepository *mocks.CityRepository
		mock           func(m *mocks.CityRepository)
		want           []*dtos.CitiesDTO
		wantErr        bool
	}{
		{
			name:           "Failed - GetCity",
			cityRepository: mocks.NewCityRepository(t),
			mock: func(m *mocks.CityRepository) {
				m.On("GetCity").Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "Success",
			cityRepository: mocks.NewCityRepository(t),
			mock: func(m *mocks.CityRepository) {
				m.On("GetCity").Return([]*dtos.CitiesDTO{}, nil)
			},
			want:    []*dtos.CitiesDTO{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &cityService{
				cityRepository: tt.cityRepository,
			}

			tt.mock(tt.cityRepository)

			got, err := s.GetAllCities()

			if (err != nil) != tt.wantErr {
				t.Errorf("cityService.GetAllCities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cityService.GetAllCities() = %v, want %v", got, tt.want)
			}
		})
	}
}
