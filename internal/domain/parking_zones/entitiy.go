package parkingzones

type ParkingZone struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Name          string  `json:"name" gorm:"type:varchar(100);not null"`
	Type          string  `json:"type" gorm:"type:varchar(50);not null;check:type IN ('general', 'ev_charging', 'covered')"`
	TotalCapacity int     `json:"total_capacity" gorm:"type:int;not null;check:total_capacity > 0"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"type:decimal(10,2);not null;check:price_per_hour > 0"`
	CreatedAt     int64   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     int64   `json:"updated_at" gorm:"autoUpdateTime"`
}
