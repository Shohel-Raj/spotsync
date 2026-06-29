package reservations

import "time"

type Reservation struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	ParkingZoneID uint      `json:"zone_id" gorm:"not null"`
	LicensePlate  string    `json:"license_plate" gorm:"type:varchar(15);not null"`
	Status        string    `json:"status" gorm:"type:varchar(20);not null;default:'active';check:status IN ('active', 'completed', 'cancelled')"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
