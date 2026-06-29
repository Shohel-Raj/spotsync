package parkingzones

import (
	"spotssync/internal/domain/parking_zones/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

// {
//   "name": "Terminal 1 EV Charging",
//   "type": "ev_charging",
//   "total_capacity": 20,
//   "price_per_hour": 5.50
// }

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
