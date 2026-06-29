package reservations

import (
	"errors"
	parkingzones "spotssync/internal/domain/parking_zones"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrorAlreadyExist = errors.New("reservation with this name already exist")
var (
	ErrParkingZoneNotFound         = errors.New("parking zone not found")
	ErrorEnoughCapacity            = errors.New("not enough capacity available")
	ErrReservationAlreadyCancelled = errors.New("reservation already cancelled")
	ErrForbiddenReservationAccess  = errors.New("you do not own this reservation")
)

type Repository interface {
	CreateReservation(reservation *Reservation) (*Reservation, error)
	GetReservationByID(id uint) (*Reservation, error)
	GetAllReservations() ([]*Reservation, error)
	UpdateReservation(id uint, reservation *Reservation) error
	DeleteReservation(id uint) error
	CancelReservation(id uint) error
	GetMyReservations(userID uint) ([]*Reservation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetReservationByID(id uint) (*Reservation, error) {
	var reservation Reservation
	result := r.db.First(&reservation, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reservation, nil
}

func (r *repository) GetAllReservations() ([]*Reservation, error) {
	var reservations []*Reservation
	result := r.db.Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func (r *repository) UpdateReservation(id uint, reservation *Reservation) error {
	result := r.db.Model(&Reservation{}).Where("id = ?", id).Updates(reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) DeleteReservation(id uint) error {
	result := r.db.Delete(&Reservation{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) CancelReservation(id uint) error {
	result := r.db.Model(&Reservation{}).Where("id = ?", id).Update("status", "canceled")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *repository) CreateReservation(ReservationData *Reservation) (*Reservation, error) {
	var booking Reservation

	// start transaction
	err := r.db.Transaction(func(tx *gorm.DB) error {

		var eventData parkingzones.ParkingZone

		err := tx.Clauses(clause.Locking{Strength: "Update"}).First(&eventData, ReservationData.ParkingZoneID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrParkingZoneNotFound
			}
			return err
		}

		if eventData.TotalCapacity < 1 {
			return ErrorEnoughCapacity
		}

		booking = Reservation{
			UserID:        ReservationData.UserID,
			ParkingZoneID: eventData.ID,
			Status:        "active",
			LicensePlate:  ReservationData.LicensePlate,
		}

		if err := tx.Create(&booking).Error; err != nil {
			return err
		}

		eventData.TotalCapacity = eventData.TotalCapacity - 1
		if err := tx.Save(&eventData).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return &booking, nil

}

func (r *repository) GetMyReservations(userID uint) ([]*Reservation, error) {
	var reservations []*Reservation

	result := r.db.Where("user_id = ?", userID).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}
