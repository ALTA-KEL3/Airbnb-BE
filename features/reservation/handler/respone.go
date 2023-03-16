package handler

import (
	// "time"

	"airbnb/features/reservation"
)

type ReservationResponseCheck struct {
	HomestayID uint   `json:"homestay_id"`
	Checkin    string `json:"checkin"`
	Checkout   string `json:"checkout"`
}

// var dateLayout1 = "2006-01-02"

func FromCoreCheck(dataCore reservation.ReservationCore) ReservationResponseCheck {
	return ReservationResponseCheck{
		HomestayID: dataCore.HomestayID,
		Checkin:    dataCore.Checkin.Format("2006-01-02"),
		Checkout:   dataCore.Checkout.Format("2006-01-02"),
	}
}

func fromCoreAvailCheck(dataCore reservation.ReservationCore) ReservationResponseCheck {
	return ReservationResponseCheck{
		HomestayID: dataCore.ID,
		Checkin:    dataCore.Checkin.Format("2006-01-02"),
		Checkout:   dataCore.Checkout.Format("2006-01-02"),
	}
}
