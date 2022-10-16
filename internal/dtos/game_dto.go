package dtos

type PlayGameRequestDTO struct {
	UserGuess string `json:"user_guess" binding:"required"`
}

type PlayGameResponseDTO struct {
	UserGuess    string `json:"user_guess"`
	CorrectGuess string `json:"correct_guess"`
	AmountWon    int    `json:"amount_won"`
}

type GetGameResponseDTO struct {
	UserID  int `json:"user_id"`
	Chance  int `json:"chance"`
	Counter int `json:"counter"`
}
