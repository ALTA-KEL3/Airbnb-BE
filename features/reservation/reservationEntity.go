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
	TotalPrice float64
	HomestayID uint
	// Homestay   homestay.Core
	UserID     uint
	// User       user.Core
}

type ReservationServiceInterface interface {
	CreateReservation(input ReservationCore) (ReservationCore, error)
}

type ReservationDataInterface interface {
	CreateReservation(input ReservationCore) (ReservationCore, error)
}
