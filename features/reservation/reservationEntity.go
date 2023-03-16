package reservation

import (
	homestay "airbnb/features/homestay"
	user "airbnb/features/user"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationCore struct {
	ID         uint
	Checkin    time.Time
	Checkout   time.Time
	TotalPrice int
	HomestayID uint
	Homestay   homestay.Core
	UserID     uint
	User       user.Core
}

type ReservationHandler interface {
	// CreateReservation() echo.HandlerFunc
	// GetReservationHistory() echo.HandlerFunc
	CheckAvailability() echo.HandlerFunc
}

type ReservationServiceInterface interface {
	CheckAvailability(input ReservationCore) (data ReservationCore, err error)
	// CreateReservation(token interface{}, totalPrice float64) (ReservationCore, string, error)
	// GetHistory(token interface{}) ([]ReservationCore, error)
}

type ReservationDataInterface interface {
	CheckAvailability(input ReservationCore) (data ReservationCore, err error)
	// CreateReservation(userID uint, totalPrice float64) (ReservationCore, string, error)
	// GetHistory(userID uint) ([]ReservationCore, error)
}
