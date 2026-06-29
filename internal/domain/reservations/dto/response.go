package dto

import "time"

type ReservationResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ParkingZoneID       uint      `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
