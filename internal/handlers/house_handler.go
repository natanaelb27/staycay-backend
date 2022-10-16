package handlers

import (
	"errors"
	"net/http"
	"strconv"

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
	querySearchCity = append(querySearchCity, c.Request.URL.Query()["searchCity"]...)
	// for _, city := range c.Request.URL.Query()["searchCity"] {
	// 	querySearchCity = append(querySearchCity, city)
	// }

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

func (h *Handler) GetHouseByIDHandler(c *gin.Context) {
	var response *dtos.GetHouseByIDResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))

	data, err := h.houseService.GetHouseByID(houseID)

	if err != nil {
		if errors.Is(err, e.ErrIdNotFoundError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetHouseByIDResponseDTO{
		HouseID:          data.HouseID,
		HouseName:        data.HouseName,
		HousePrice:       data.HousePrice,
		HouseDescription: data.HouseDescription,
		HouseCityName:    data.HouseCityName,
		HouseMaxGuest:    data.HouseMaxGuest,
		HouseStatus:      data.HouseStatus,
		HousePhoto:       data.HousePhoto,
		HostID:           data.HostID,
		HostName:         data.HostName,
		HostEmail:        data.HostEmail,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetHostHousesHandler(c *gin.Context) {
	var response *dtos.GetHostHousesResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	data, err := h.houseService.GetHostHouses(userContext.UserID)
	if err != nil {
		if errors.Is(err, e.ErrInvalidRoleNotHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotHostError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetHostHousesResponseDTO{
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
		if errors.Is(err, e.ErrInvalidRoleNotHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotHostError.Error()))
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

func (h *Handler) UpdateHouseHandler(c *gin.Context) {
	var request *dtos.UpdateHouseRequestDTO
	var response *dtos.UpdateHouseResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))

	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	newHouse, err := h.houseService.UpdateHouse(houseID, userContext.UserID, request)

	if err != nil {
		if errors.Is(err, e.ErrInvalidRoleNotHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotHostError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidHostError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.UpdateHouseResponseDTO{
		Name:        newHouse.Name,
		Price:       newHouse.Price,
		Description: newHouse.Description,
		CityID:      newHouse.CityID,
		MaxGuest:    newHouse.MaxGuest,
		Status:      newHouse.Status,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}

func (h *Handler) DeleteHouseHandler(c *gin.Context) {
	var response *dtos.DeleteHouseResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))
	userContext := c.MustGet("user").(dtos.JwtData)

	data, err := h.houseService.DeleteHouse(houseID, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrIdNotFoundError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidHostError.Error()))
			return
		}

		if errors.Is(err, e.ErrDeleteUpcomingReservationError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrDeleteUpcomingReservationError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.DeleteHouseResponseDTO{
		HouseID:     data.ID,
		Price:       data.Price,
		Description: data.Description,
		CityID:      data.CityID,
		MaxGuest:    data.MaxGuest,
		Status:      data.Status,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}

func (h *Handler) AddHousePhotoHandler(c *gin.Context) {
	var request *dtos.AddHousePhotoRequestDTO
	var response *dtos.AddHousePhotoResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))

	userContext := c.MustGet("user").(dtos.JwtData)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	photoList, err := h.houseService.AddHousePhoto(houseID, userContext.UserID, request)

	if err != nil {
		if errors.Is(err, e.ErrInvalidRoleNotHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotHostError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidHostError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.AddHousePhotoResponseDTO{
		Photo: photoList,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}

func (h *Handler) DeleteHousePhotoHandler(c *gin.Context) {
	var response *dtos.DeleteHousePhotoResponseDTO

	housePhotoID, _ := strconv.Atoi(c.Param("housePhotoId"))
	userContext := c.MustGet("user").(dtos.JwtData)

	data, err := h.houseService.DeleteHousePhoto(housePhotoID, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrIdNotFoundError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidHostError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.DeleteHousePhotoResponseDTO{
		HouseID: data.HouseID,
		Photo:   data.Photo,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}
