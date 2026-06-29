package reservations

import (
	"github.com/labstack/echo/v5"
	"spotssync/internal/auth"
	"spotssync/internal/config"
	"spotssync/internal/middlewares"

	"gorm.io/gorm"
)

func ReservationRegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {

	reservationRepository := NewRepository(db)
	reservationService := NewService(reservationRepository)
	reservationHandler := NewHandler(reservationService)
	jwtService := auth.NewJWTService(cfg.JwtSecret)

	api := e.Group("/api/v1/reservations")

	api.POST("", reservationHandler.CreateReservation, middlewares.AuthMiddleware(jwtService))

	api.GET("/my-reservations", reservationHandler.GetMyReservations, middlewares.AuthMiddleware(jwtService))

	api.GET("", reservationHandler.GetReservations, middlewares.AuthMiddleware(jwtService), middlewares.AdminOnly())
	api.GET("/:id", reservationHandler.GetReservationByID, middlewares.AuthMiddleware(jwtService))
	// api.DELETE("/:id", reservationHandler.DeleteReservation, middlewares.AuthMiddleware(jwtService), middlewares.AdminOnly())

	api.DELETE("/:id", reservationHandler.DeleteReservation)

}
