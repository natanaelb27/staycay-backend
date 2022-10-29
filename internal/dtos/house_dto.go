package dtos

import "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"

// type HouseDetailsDTO struct {
// 	HouseID          int            `json:"house_id"`
// 	HouseName        string         `json:"house_name"`
// 	HousePrice       int            `json:"house_price"`
// 	HouseDescription string         `json:"description"`
// 	HouseCityName    string         `json:"house_city_name"`
// 	HouseMaxGuest    int            `json:"house_max_guest"`
// 	HouseStatus      string         `json:"house_status"`
// 	HousePhoto       pq.StringArray `json:"house_photo" gorm:"type:text[]"`
// 	HostID           int            `json:"host_id"`
// 	HostName         string         `json:"host_name"`
// 	HostEmail        string         `json:"host_email"`
// }

type HousePhotosDTO struct {
	PhotoID int    `json:"photo_id"`
	Photo   string `json:"photo"`
}

type HouseDetailsDTO struct {
	HouseID          int    `json:"house_id"`
	HouseName        string `json:"house_name"`
	HousePrice       int    `json:"house_price"`
	HouseDescription string `json:"house_description"`
	HouseCityID      int    `json:"house_city_id"`
	HouseCityName    string `json:"house_city_name"`
	HouseMaxGuest    int    `json:"house_max_guest"`
	HouseStatus      string `json:"house_status"`
	HostID           int    `json:"host_id"`
	HostName         string `json:"host_name"`
	HostEmail        string `json:"host_email"`
}

type HouseDetailsWithPhotoDTO struct {
	HouseID          int               `json:"house_id"`
	HouseName        string            `json:"house_name"`
	HousePrice       int               `json:"house_price"`
	HouseDescription string            `json:"house_description"`
	HouseCityID      int               `json:"house_city_id"`
	HouseCityName    string            `json:"house_city_name"`
	HouseMaxGuest    int               `json:"house_max_guest"`
	HouseStatus      string            `json:"house_status"`
	HousePhoto       []*HousePhotosDTO `json:"house_photo"`
	HostID           int               `json:"host_id"`
	HostName         string            `json:"host_name"`
	HostEmail        string            `json:"host_email"`
}

type GetAllHousesResponseDTO struct {
	Pagination *helpers.Pagination         `json:"pagination"`
	House      []*HouseDetailsWithPhotoDTO `json:"house"`
}

type GetHouseByIDResponseDTO struct {
	HouseID          int               `json:"house_id"`
	HouseName        string            `json:"house_name"`
	HousePrice       int               `json:"house_price"`
	HouseDescription string            `json:"house_description"`
	HouseCityID      int               `json:"house_city_id"`
	HouseCityName    string            `json:"house_city_name"`
	HouseMaxGuest    int               `json:"house_max_guest"`
	HouseStatus      string            `json:"house_status"`
	HousePhoto       []*HousePhotosDTO `json:"house_photo"`
	HostID           int               `json:"host_id"`
	HostName         string            `json:"host_name"`
	HostEmail        string            `json:"host_email"`
}

type GetHostHousesResponseDTO struct {
	Pagination *helpers.Pagination         `json:"pagination"`
	House      []*HouseDetailsWithPhotoDTO `json:"house"`
}

type HouseStatusDTO struct {
	HouseStatusID   int    `json:"house_status_id"`
	HouseStatusName string `json:"house_status_name"`
}

type GetAllHouseStatusesResponseDTO struct {
	HouseStatus []*HouseStatusDTO `json:"house_status"`
}

type AddNewHouseRequestDTO struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	CityID      int    `json:"city_id" binding:"required"`
	// CityName    string   `json:"city_name" binding:"required"`
	MaxGuest int      `json:"max_guest" binding:"required"`
	Photo    []string `json:"photo" binding:"required"`
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
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	CityID      int    `json:"city_id"`
	// CityName    string `json:"city_name"`
	MaxGuest      int `json:"max_guest"`
	HouseStatusID int `json:"house_status_id"`
	// Status   string `json:"status"`
}

type UpdateHouseResponseDTO struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	CityID        int    `json:"city_id"`
	MaxGuest      int    `json:"max_guest"`
	HouseStatusID int    `json:"house_status_id"`
	// Status      string `json:"status"`
}

type AddHousePhotoRequestDTO struct {
	Photo []string `json:"photo"`
}

type AddHousePhotoResponseDTO struct {
	Photo []string `json:"photo"`
}

type DeleteHouseResponseDTO struct {
	HouseID       int    `json:"house_id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	CityID        int    `json:"city_id"`
	MaxGuest      int    `json:"max_guest"`
	HouseStatusID int    `json:"house_status_id"`
	// Status      string `json:"status"`
}

type DeleteHousePhotoResponseDTO struct {
	HouseID int    `json:"house_id"`
	Photo   string `json:"photo"`
}
