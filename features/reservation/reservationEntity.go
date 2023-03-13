package reservation

import (
	// "airbnb/features/homestay"
	// "airbnb/features/user"
	"time"
)

type ReservationCore struct {
	ID uint
	Checkin    time.Time
	Checkout   time.Time
	TotalPrice int
	HomestayID uint
	// Homestay   homestay.Core
	UserID     uint
	// User       user.Core
}

type ReservationServiceInterface interface {
	// GetAll() ([]ReservationCore, error)
	// GetById(id uint) (ReservationCore, error)
	CreateReservation(input ReservationCore) error
	// Update(reservationCore ReservationCore, id uint) (ReservationCore, error)
	// Delete(id uint) error
}

type ReservationDataInterface interface {
	// SelectAll() ([]ReservationCore, error)
	// SelectById(id uint) (ReservationCore, error)
	CreateReservation(input ReservationCore) error
	// Edit(reservationCore ReservationCore, id uint) error
	// Destroy(id uint) error
}
