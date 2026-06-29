package parkingzones

import (
	"errors"
	"net/http"
	"spotssync/internal/domain/parking_zones/dto"
	"spotssync/internal/httpresponse"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
	"strconv"
)

type handler struct {
	service *service
}

func parkingZoneErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Parking Zone not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}
func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateParkingZone(c *echo.Context) error {
	var req dto.CreateParkingZoneRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.Name == "" || req.Type == "" || req.TotalCapacity <= 0 || req.PricePerHour < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	response, err := h.service.CreateParkingZone(req)
	if err != nil {
		if errors.Is(err, ErrorAlreadyExist) {
			return c.JSON(http.StatusConflict, httpresponse.Error{
				Code:    http.StatusConflict,
				Message: "Failed to create Parking Zone",
				Details: err.Error(),
			})
		}
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Parking Zone created successfully",
		"data":    response,
	})
}

func (h *handler) GetParkingZones(c *echo.Context) error {
	parkingZones, err := h.service.GetParkingZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve Parking Zones",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking Zones retrieved successfully",
		"data":    parkingZones,
	})
}

func (h *handler) GetParkingZoneByID(c *echo.Context) error {
	parkingZoneID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Parking Zone ID",
			Details: err.Error(),
		})
	}
	parkingZone, err := h.service.GetParkingZoneByID(uint(parkingZoneID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, httpresponse.Error{
				Code:    http.StatusNotFound,
				Message: "Parking Zone not found",
				Details: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve Parking Zone",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking Zone retrieved successfully",
		"data":    parkingZone,
	})
}

func (h *handler) updateParkingZone(c *echo.Context) error {

	parkingZoneID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Parking Zone ID",
			Details: err.Error(),
		})
	}
	var req dto.UpdateParkingZoneRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.UpdateParkingZone(uint(parkingZoneID), req)
	if err != nil {
		return parkingZoneErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking Zone updated successfully",
		"data":    response,
	})

}

func (h *handler) DeleteParkingZone(c *echo.Context) error {
	parkingZoneID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid Parking Zone ID",
			Details: err.Error(),
		})
	}
	parkingZone, err := h.service.GetParkingZoneByID(uint(parkingZoneID))
	if err != nil || parkingZone == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, httpresponse.Error{
				Code:    http.StatusNotFound,
				Message: "Parking Zone not found",
				Details: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve Parking Zone",
			Details: err.Error(),
		})
	}

	err = h.service.DeleteParkingZone(uint(parkingZoneID))
	if err != nil {
		return parkingZoneErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Parking Zone deleted successfully",
	})
}
