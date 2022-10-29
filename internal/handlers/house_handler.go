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

	data, pagination, err := h.houseService.GetAllHouses(queryParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetAllHousesResponseDTO{
		Pagination: pagination,
		House:      data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetHouseByIDHandler(c *gin.Context) {
	var response *dtos.GetHouseByIDResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))

	queryParams := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = values[0]
	}

	data, err := h.houseService.GetHouseByID(houseID, queryParams)

	if err != nil {
		if errors.Is(err, e.ErrHouseNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrHouseNotFoundError.Error()))
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
		HouseCityID:      data.HouseCityID,
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

// func (h *Handler) GetUnscopedHouseByIDHandler(c *gin.Context) {
// 	var response *dtos.GetHouseByIDResponseDTO

// 	houseID, _ := strconv.Atoi(c.Param("houseId"))

// 	data, err := h.houseService.GetUnscopedHouseByID(houseID)

// 	if err != nil {
// 		if errors.Is(err, e.ErrIdNotFoundError) {
// 			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrIdNotFoundError.Error()))
// 			return
// 		}

// 		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	response = &dtos.GetHouseByIDResponseDTO{
// 		HouseID:          data.HouseID,
// 		HouseName:        data.HouseName,
// 		HousePrice:       data.HousePrice,
// 		HouseDescription: data.HouseDescription,
// 		HouseCityID:      data.HouseCityID,
// 		HouseCityName:    data.HouseCityName,
// 		HouseMaxGuest:    data.HouseMaxGuest,
// 		HouseStatus:      data.HouseStatus,
// 		HousePhoto:       data.HousePhoto,
// 		HostID:           data.HostID,
// 		HostName:         data.HostName,
// 		HostEmail:        data.HostEmail,
// 	}

// 	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
// }

func (h *Handler) GetHostHousesHandler(c *gin.Context) {
	var response *dtos.GetHostHousesResponseDTO

	userContext := c.MustGet("user").(dtos.JwtData)

	queryParams := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = values[0]
	}

	data, pagination, err := h.houseService.GetHostHouses(userContext.UserID, queryParams)
	if err != nil {
		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidRoleNotHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotHostError.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetHostHousesResponseDTO{
		Pagination: pagination,
		House:      data,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}

func (h *Handler) GetAllHouseStatusesHandler(c *gin.Context) {
	var response *dtos.GetAllHouseStatusesResponseDTO

	data, err := h.houseService.GetAllHouseStatuses()

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = &dtos.GetAllHouseStatusesResponseDTO{
		HouseStatus: data,
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
		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

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
		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrHouseNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrHouseNotFoundError.Error()))
			return
		}

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
		Name:          newHouse.Name,
		Price:         newHouse.Price,
		Description:   newHouse.Description,
		CityID:        newHouse.CityID,
		MaxGuest:      newHouse.MaxGuest,
		HouseStatusID: newHouse.HouseStatusID,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))

}

func (h *Handler) DeleteHouseHandler(c *gin.Context) {
	var response *dtos.DeleteHouseResponseDTO

	houseID, _ := strconv.Atoi(c.Param("houseId"))
	userContext := c.MustGet("user").(dtos.JwtData)

	data, err := h.houseService.DeleteHouse(houseID, userContext.UserID)

	if err != nil {
		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrHouseNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrHouseNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrInvalidRoleNotHostError) {
			c.JSON(http.StatusForbidden, helpers.BuildErrorResponse(http.StatusForbidden, e.ErrInvalidRoleNotHostError.Error()))
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
		HouseID:       data.ID,
		Name:          data.Name,
		Price:         data.Price,
		Description:   data.Description,
		CityID:        data.CityID,
		MaxGuest:      data.MaxGuest,
		HouseStatusID: data.HouseStatusID,
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
		if errors.Is(err, e.ErrUserNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrUserNotFoundError.Error()))
			return
		}

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
		if errors.Is(err, e.ErrHousePhotoNotFoundError) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, e.ErrHousePhotoNotFoundError.Error()))
			return
		}

		if errors.Is(err, e.ErrHousePhotoMustMoreThanOneError) {
			c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, e.ErrHousePhotoMustMoreThanOneError.Error()))
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
