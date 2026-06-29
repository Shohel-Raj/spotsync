package parkingzones

import (
	"spotssync/internal/domain/parking_zones/dto"

	"gorm.io/gorm"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateParkingZone(req dto.CreateParkingZoneRequest) (*dto.ParkingZoneResponse, error) {

	parkingZone := ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	err := s.repo.CreateParkingZone(&parkingZone)
	if err != nil {
		return nil, err
	}
	return &dto.ParkingZoneResponse{
		ID:            parkingZone.ID,
		Name:          parkingZone.Name,
		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
	}, nil
}

func (s *service) GetParkingZones() ([]dto.ParkingZoneResponse, error) {
	parkingZones, err := s.repo.GetAllParkingZones()
	if err != nil {
		return nil, err
	}
	var responses []dto.ParkingZoneResponse
	for _, zone := range parkingZones {
		responses = append(responses, dto.ParkingZoneResponse{
			ID:            zone.ID,
			Name:          zone.Name,
			Type:          zone.Type,
			TotalCapacity: zone.TotalCapacity,
			PricePerHour:  zone.PricePerHour,
		})
	}
	return responses, nil
}

func (s *service) GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error) {
	parkingZone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.ParkingZoneResponse{
		ID:            parkingZone.ID,
		Name:          parkingZone.Name,
		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
	}, nil
}

func (s *service) UpdateParkingZone(id uint, req dto.UpdateParkingZoneRequest) (*dto.ParkingZoneResponse, error) {
	parkingZone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		parkingZone.Name = *req.Name
	}
	if req.Type != nil {
		parkingZone.Type = *req.Type
	}
	if req.TotalCapacity != nil {
		parkingZone.TotalCapacity = *req.TotalCapacity
	}
	if req.PricePerHour != nil {
		parkingZone.PricePerHour = *req.PricePerHour
	}
	if parkingZone == nil {
		return nil, gorm.ErrRecordNotFound
	}
	err = s.repo.UpdateParkingZone(id, parkingZone)
	if err != nil {
		return nil, err
	}
	return &dto.ParkingZoneResponse{
		ID:            parkingZone.ID,
		Name:          parkingZone.Name,
		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
	}, nil
}

func (s *service) DeleteParkingZone(id uint) error {
	err := s.repo.DeleteParkingZone(id)
	if err != nil {
		return err
	}
	return nil
}
