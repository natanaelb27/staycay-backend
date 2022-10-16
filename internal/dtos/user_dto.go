package dtos

type UserDetailsDTO struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
	Role     string `json:"role"`
	WalletID int    `json:"wallet_id"`
	Balance  int    `json:"balance"`
}

type GetUserProfileResponseDTO struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
	Role     string `json:"role"`
	WalletID int    `json:"wallet_id"`
	Balance  int    `json:"balance"`
}

type GetUserProfileByUserIDResponseDTO struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
	Role     string `json:"role"`
	WalletID int    `json:"wallet_id"`
	Balance  int    `json:"balance"`
}

type UpdateUserProfileRequestDTO struct {
	Name        string `json:"name"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	CityID      int    `json:"city_id"`
	// CityName    string `json:"city_name"`
}

type UpdateUserProfileResponseDTO struct {
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	CityID  int    `json:"city_id"`
}

type UpdateUserRoleResponseDTO struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}
