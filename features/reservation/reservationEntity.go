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
	TotalPrice float64
	HomestayID uint
	// Homestay   homestay.Core
	UserID     uint
	// User       user.Core
}

type ReservationHandler interface {
	CreateReservation() echo.HandlerFunc
	GetReservationHistory() echo.HandlerFunc
}

type ReservationServiceInterface interface {
	CreateReservation(token interface{}, totalPrice float64) (ReservationCore, string, error)
	GetOrderHistory(token interface{}) ([]ReservationCore, error)
}

type ReservationDataInterface interface {
	CreateReservation(userID uint, totalPrice float64) (ReservationCore, string, error)
	GetOrderHistory(userID uint) ([]ReservationCore, error)
}
