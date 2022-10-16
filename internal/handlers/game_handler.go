package handlers

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) PlayGameHandler(c *gin.Context) {
	var request *dtos.PlayGameRequestDTO
	var response *dtos.PlayGameResponseDTO
	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	correctGuess, amountWon, err := h.gameService.PlayGame(request.UserGuess, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrInsufficientGameChanceError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInsufficientGameChanceError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.PlayGameResponseDTO{
		UserGuess:    request.UserGuess,
		CorrectGuess: correctGuess,
		AmountWon:    amountWon,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetGameHandler(c *gin.Context) {
	var response *dtos.GetGameResponseDTO
	userContext := c.MustGet("user").(dtos.JwtData)

	gameData, err := h.gameService.GetGame(userContext.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetGameResponseDTO{
		UserID:  gameData.UserID,
		Chance:  gameData.Chance,
		Counter: gameData.Counter,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
