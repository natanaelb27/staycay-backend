package dtos

type RegisterRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	CityID   int    `json:"city_id" binding:"required"`
	// CityName string `json:"city_name" binding:"required"`
}

type RegisterResponseDTO struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type JwtData struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}
