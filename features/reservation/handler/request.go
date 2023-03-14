package handler

import (
	"airbnb/features/reservation"
	"time"
)

type ReservationRequest struct {
	Checkin    string  `json:"checkin" form:"checkin"`
	Checkout   string  `json:"checkout" form:"checkout"`
	TotalPrice float64 `json:"total_price " form:"total_price"`
	HomestayID uint    `json:"homestay_id" form:"homestay_id"`
	UserID     uint    `json:"user_id" form:"user_id"`
}

func RequestToCore(data ReservationRequest, Checkin, Checkout time.Time) reservation.ReservationCore {
	return reservation.ReservationCore{
		HomestayID: data.HomestayID,
		TotalPrice: data.TotalPrice,
		UserID:     data.UserID,
		Checkin:    Checkin,
		Checkout:   Checkout,
	}
}
