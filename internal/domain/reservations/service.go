package reservations

import (
	"spotssync/internal/domain/reservations/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateReservation(req dto.CreateReservationRequest) (*dto.ReservationResponse, error) {
	reservation := Reservation{
		UserID:        req.UserID,
		LicensePlate:  req.LicensePlate,
		ParkingZoneID: req.ParkingZoneID,
		Status:        "active",
	}
	created, err := s.repo.CreateReservation(&reservation)
	if err != nil {
		return nil, err
	}
	return &dto.ReservationResponse{
		ID:            created.ID,
		UserID:        created.UserID,
		LicensePlate:  created.LicensePlate,
		ParkingZoneID: created.ParkingZoneID,
		Status:        created.Status,
	}, nil
}

func (s *service) GetReservationByID(id uint) (*dto.ReservationResponse, error) {
	reservation, err := s.repo.GetReservationByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.ReservationResponse{
		ID:            reservation.ID,
		UserID:        reservation.UserID,
		LicensePlate:  reservation.LicensePlate,
		ParkingZoneID: reservation.ParkingZoneID,
		Status:        reservation.Status,
		CreatedAt:     reservation.CreatedAt,
		UpdatedAt:     reservation.UpdatedAt,
	}, nil
}

func (s *service) UpdateReservationStatus(id uint, status string) (*dto.ReservationResponse, error) {
	reservation, err := s.repo.GetReservationByID(id)
	if err != nil || reservation == nil {
		return nil, err
	}
	reservation.Status = status
	err = s.repo.UpdateReservation(id, reservation)
	if err != nil {
		return nil, err
	}
	return &dto.ReservationResponse{
		ID:            reservation.ID,
		UserID:        reservation.UserID,
		LicensePlate:  reservation.LicensePlate,
		ParkingZoneID: reservation.ParkingZoneID,
		Status:        reservation.Status,
		CreatedAt:     reservation.CreatedAt,
		UpdatedAt:     reservation.UpdatedAt,
	}, nil
}

func (s *service) CancelReservation(id uint) (*dto.ReservationResponse, error) {
	reservation, err := s.repo.GetReservationByID(id)
	if err != nil || reservation == nil {
		return nil, err
	}
	reservation.Status = "cancelled"
	err = s.repo.UpdateReservation(id, reservation)
	if err != nil {
		return nil, err
	}
	return &dto.ReservationResponse{
		ID:            reservation.ID,
		UserID:        reservation.UserID,
		LicensePlate:  reservation.LicensePlate,
		ParkingZoneID: reservation.ParkingZoneID,
		Status:        reservation.Status,
		CreatedAt:     reservation.CreatedAt,
		UpdatedAt:     reservation.UpdatedAt,
	}, nil
}
func (s *service) GetAllReservations() ([]*dto.ReservationResponse, error) {
	reservations, err := s.repo.GetAllReservations()
	if err != nil {
		return nil, err
	}
	var responses []*dto.ReservationResponse
	for _, reservation := range reservations {
		responses = append(responses, &dto.ReservationResponse{
			ID:            reservation.ID,
			UserID:        reservation.UserID,
			LicensePlate:  reservation.LicensePlate,
			ParkingZoneID: reservation.ParkingZoneID,
			Status:        reservation.Status,
			CreatedAt:     reservation.CreatedAt,
			UpdatedAt:     reservation.UpdatedAt,
		})
	}
	return responses, nil
}
