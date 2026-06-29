package parkingzones

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("parking zone with this name already exist")

type Repository interface {
	CreateParkingZone(zone *ParkingZone) error
	GetParkingZoneByID(id uint) (*ParkingZone, error)
	GetAllParkingZones() ([]*ParkingZone, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateParkingZone(zone *ParkingZone) error {
	result := r.db.Create(zone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r repository) GetParkingZoneByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone
	result := r.db.First(&zone, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &zone, nil
}

func (r repository) GetAllParkingZones() ([]*ParkingZone, error) {
	var zones []*ParkingZone
	result := r.db.Find(&zones)
	if result.Error != nil {
		return nil, result.Error
	}
	return zones, nil
}
