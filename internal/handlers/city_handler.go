package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelb27/staycay-backend/internal/dtos"
	"github.com/natanaelb27/staycay-backend/internal/helpers"
)

func (h *Handler) GetAllCitiesHandler(c *gin.Context) {
	var response *dtos.GetAllCitiesResponseDTO

	data, err := h.cityService.GetAllCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetAllCitiesResponseDTO{
		City: data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
