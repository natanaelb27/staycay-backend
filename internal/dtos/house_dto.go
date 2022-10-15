package dtos

import "github.com/lib/pq"

type HouseDetailsDTO struct {
	HouseID          int            `json:"house_id"`
	HouseName        string         `json:"house_name"`
	HousePrice       int            `json:"house_price"`
	HouseDescription string         `json:"description"`
	HouseCityName    string         `json:"house_city_name"`
	HouseMaxGuest    int            `json:"house_max_guest"`
	HouseStatus      string         `json:"house_status"`
	HousePhoto       pq.StringArray `json:"house_photo" gorm:"type:text[]"`
	HostID           int            `json:"host_id"`
	HostName         string         `json:"host_name"`
	HostEmail        string         `json:"host_email"`
}

type GetAllHousesResponseDTO struct {
	House []*HouseDetailsDTO `json:"house"`
}

type GetHostHousesResponseDTO struct {
	House []*HouseDetailsDTO `json:"house"`
}

type AddNewHouseRequestDTO struct {
	Name        string   `json:"name" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Description string   `json:"description" binding:"required"`
	CityName    string   `json:"city_name" binding:"required"`
	MaxGuest    int      `json:"max_guest" binding:"required"`
	Photo       []string `json:"photo" binding:"required"`
}

type AddNewHouseResponseDTO struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	CityID      int      `json:"city_id"`
	MaxGuest    int      `json:"max_guest"`
	Photo       []string `json:"photo"`
}

// type UpdateHouseRequestDTO struct {
// 	Name        string   `json:"name" binding:"required"`
// 	Price       int      `json:"price" binding:"required"`
// 	Description string   `json:"description" binding:"required"`
// 	CityName    string   `json:"city_name" binding:"required"`
// 	MaxGuest    int      `json:"max_guest" binding:"required"`
// 	Photo       []string `json:"photo"`
// 	Status      string   `json:"status" binding:"required"`
// }

type UpdateHouseRequestDTO struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	CityName    string   `json:"city_name"`
	MaxGuest    int      `json:"max_guest"`
	Photo       []string `json:"photo"`
	Status      string   `json:"status"`
}

type UpdateHouseResponseDTO struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	CityID      int      `json:"city_id"`
	MaxGuest    int      `json:"max_guest"`
	Photo       []string `json:"photo"`
	Status      string   `json:"status"`
}
