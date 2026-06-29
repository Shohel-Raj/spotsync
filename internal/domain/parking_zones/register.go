package parkingzones

import (
	"github.com/labstack/echo/v5"
	"spotssync/internal/auth"
	"spotssync/internal/config"
	"spotssync/internal/middlewares"

	"gorm.io/gorm"
)

func ParkingZoneRegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {

	parkingRepository := NewRepository(db)
	parkingService := NewService(parkingRepository)
	parkingHandler := NewHandler(parkingService)
	jwtService := auth.NewJWTService(cfg.JwtSecret)

	api := e.Group("/api/v1/parkingzones")

	api.POST("", parkingHandler.CreateParkingZone, middlewares.AuthMiddleware(jwtService), middlewares.AdminOnly())
	api.GET("", parkingHandler.GetParkingZones)
	api.GET("/:id", parkingHandler.GetParkingZoneByID)
	api.PUT("/:id", parkingHandler.updateParkingZone, middlewares.AuthMiddleware(jwtService), middlewares.AdminOnly())
	api.DELETE("/:id", parkingHandler.DeleteParkingZone, middlewares.AuthMiddleware(jwtService), middlewares.AdminOnly())

}
