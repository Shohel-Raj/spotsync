package parkingzones

import (
	"spotssync/internal/config"

	"github.com/labstack/echo/v5"

	"gorm.io/gorm"
)

func ParkingZoneRegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {

	parlkingRepository := NewRepository(db)
	parkingService := NewService(parlkingRepository)
	parkingHandler := NewHandler(parkingService)

	api := e.Group("/api/v1/parkingzones")

	api.POST("", parkingHandler.CreateParkingZone)
	api.GET("", parkingHandler.GetParkingZones)
	api.GET("/:id", parkingHandler.GetParkingZoneByID)

}
