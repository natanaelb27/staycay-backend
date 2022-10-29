package handlers

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) TopUpWalletHandler(c *gin.Context) {
	var request *dtos.TopUpWalletRequestDTO
	var response *dtos.TopUpWalletResponseDTO
	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	data, err := h.walletService.TopUpWallet(userContext.UserID, request.Amount)

	if err != nil {
		if errors.Is(err, e.ErrTopUpTooSmallError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrTopUpTooSmallError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.TopUpWalletResponseDTO{
		UserID:   data.UserID,
		WalletID: data.ID,
		Balance:  data.Balance,
		Amount:   request.Amount,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
