package handlers

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	var request dtos.RegisterRequestDTO
	var response dtos.RegisterResponseDTO

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	newAuth, token, err := h.authService.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = dtos.RegisterResponseDTO{
		ID:    newAuth.ID,
		Email: newAuth.Email,
		Token: *token,
	}

	c.JSON(http.StatusCreated, helpers.BuildSuccessResponse(http.StatusCreated, http.StatusText(http.StatusCreated), response))
}
