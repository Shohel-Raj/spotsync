package reservations

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"spotssync/internal/domain/reservations/dto"
	"spotssync/internal/httpresponse"

	"github.com/labstack/echo/v5"
	"strconv"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func reservationErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Success: false,
			Code:    http.StatusNotFound,
			Message: "Reservation not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Success: false,
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateReservation(c *echo.Context) error {
	var req dto.CreateReservationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Invalid parking zone type",
			Details: err.Error(),
		})
	}
	userIDRaw := c.Get("user_id")
	if userIDRaw == nil {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	var userID uint
	switch v := userIDRaw.(type) {
	case uint:
		userID = v
	case float64:
		userID = uint(v) // ✅ JWT claims come as float64
	case int:
		userID = uint(v)
	default:
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid user ID in token",
		})
	}

	reservation, err := h.service.CreateReservation(req, userID)
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success":     true,
		"message":     "Reservation confirm successfully",
		"reservation": reservation,
	})
}

func (h *handler) GetReservations(c *echo.Context) error {
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":      true,
		"message":      "Reservations retrieved successfully",
		"reservations": reservations,
	})
}

func (h *handler) GetReservationByID(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	reservation, err := h.service.GetReservationByID(uint(id))
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":     true,
		"message":     "Reservation retrieved successfully",
		"reservation": reservation,
	})
}

func (h *handler) DeleteReservation(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return reservationErrorResponse(c, err)
	}
	reservation, err := h.service.GetReservationByID(uint(id))
	if err != nil || reservation == nil {
		return reservationErrorResponse(c, err)
	}

	if err := h.service.DeleteReservation(uint(id)); err != nil {
		return reservationErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Reservation cancelled successfully",
	})
}

func (h *handler) GetMyReservations(c *echo.Context) error {
	userIDRaw := c.Get("user_id")
	var userID uint
	switch v := userIDRaw.(type) {
	case uint:
		userID = v
	case float64:
		userID = uint(v)
	default:
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	reservations, err := h.service.GetMyReservations(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve reservations",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":      true,
		"message":      "Reservations retrieved successfully",
		"reservations": reservations,
	})
}

func (h *handler) CancelReservation(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid reservation ID",
			Details: err.Error(),
		})
	}

	result, err := h.service.CancelReservation(uint(id))
	if err != nil {
		return reservationErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Reservation cancelled successfully",
		"data":    result,
	})

}
