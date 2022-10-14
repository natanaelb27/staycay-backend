package dtos

type CitiesDTO struct {
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

type GetAllCitiesResponseDTO struct {
	City []*CitiesDTO `json:"city"`
}
