package dtos

type HouseDetailsDTO struct {
	HouseID     int    `json:"house_id"`
	HouseName   string `json:"house_name"`
	HousePrice  int    `json:"house_price"`
	Description string `json:"description"`
	CityName    string `json:"city_name"`
	MaxGuest    int    `json:"max_guest"`
	HostName    string `json:"host_name"`
	HostEmail   string `json:"host_email"`
}

type GetAllHousesResponseDTO struct {
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
