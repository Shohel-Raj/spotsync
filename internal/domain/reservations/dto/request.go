package dto

type CreateReservationRequest struct {
	ParkingZoneID uint   `json:"zone_id"`
	LicensePlate  string `json:"license_plate"`
}
