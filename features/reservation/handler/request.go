package handler

import (
	"time"

	"airbnb/features/reservation"
)

type ReservationRequestCheck struct {
	Checkin    string `json:"checkin" form:"checkin"`
	Checkout   string `json:"checkout" form:"checkout"`
	HomestayID int    `json:"homestay_id" form:"homestay_id"`
}

var dateLayout = "2006-01-02"


func ToCoreCheck(reservationInput ReservationRequestCheck) reservation.ReservationCore {
	in, _ := time.Parse(dateLayout, reservationInput.Checkin)
	out, _ := time.Parse(dateLayout, reservationInput.Checkout)
	return reservation.ReservationCore{
		HomestayID: uint(reservationInput.HomestayID),
		Checkin:    in,
		Checkout:   out,
	}
}

