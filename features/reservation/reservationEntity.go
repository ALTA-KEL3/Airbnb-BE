package reservation

import (
	// "airbnb/features/homestay"
	// "airbnb/features/user"
	"github.com/labstack/echo/v4"
	"time"
)

type ReservationCore struct {
	ID uint
	Checkin    time.Time
	Checkout   time.Time
	Duration   int
	TotalPrice int
	HomestayID uint
	// Homestay   homestay.Core
	Homestay Homestay
	UserID     uint
	// User       user.Core
}

type Homestay struct {
	ID            uint
	Name          string
	Address       string
	Price int
	BookedStart   time.Time
	BookedEnd     time.Time
	Reservation   []ReservationCore
}


type ReservationHandler interface {
	// CreateReservation() echo.HandlerFunc
	// GetReservationHistory() echo.HandlerFunc
	CheckAvailability() echo.HandlerFunc
}

type ReservationServiceInterface interface {
	CheckAvailability(input ReservationCore) (data Homestay, err error)
	// CreateReservation(token interface{}, totalPrice float64) (ReservationCore, string, error)
	// GetHistory(token interface{}) ([]ReservationCore, error)
}

type ReservationDataInterface interface {
	CheckAvailability(input ReservationCore) (data Homestay, err error)
	// CreateReservation(userID uint, totalPrice float64) (ReservationCore, string, error)
	// GetHistory(userID uint) ([]ReservationCore, error)
}
