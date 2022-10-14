package dtos

type TopUpWalletRequestDTO struct {
	Amount int `json:"amount" binding:"required"`
}

type TopUpWalletResponseDTO struct {
	UserID   int `json:"user_id"`
	WalletID int `json:"wallet_id"`
	Balance  int `json:"balance"`
	Amount   int `json:"amount"`
}
