package dto

type CreateReservationRequest struct {
	UserID        uint   `json:"user_id"`
	ParkingZoneID uint   `json:"zone_id"`
	LicensePlate  string `json:"license_plate"`
}
