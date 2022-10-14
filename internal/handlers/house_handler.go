package handlers

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllHousesHandler(c *gin.Context) {
	var response *dtos.GetAllHousesResponseDTO

	queryParams := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = values[0]
	}

	querySearchCity := make([]string, 0)
	for _, city := range c.Request.URL.Query()["searchCity"] {
		querySearchCity = append(querySearchCity, city)
	}

	data, err := h.houseService.GetAllHouses(queryParams, querySearchCity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetAllHousesResponseDTO{
		House: data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) AddNewHouseHandler(c *gin.Context) {
	var request *dtos.AddNewHouseRequestDTO
	var response *dtos.AddNewHouseResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	newHouse, photo, err := h.houseService.AddNewHouse(userContext.UserID, request)
	if err != nil {
		if errors.Is(err, e.ErrInvalidRoleAddHouseError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrInvalidRoleAddHouseError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.AddNewHouseResponseDTO{
		Name:        newHouse.Name,
		Price:       newHouse.Price,
		Description: newHouse.Description,
		CityID:      newHouse.CityID,
		MaxGuest:    newHouse.MaxGuest,
		Photo:       photo,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}
