package handler

import (
	// "time"

	"airbnb/features/reservation"
)

type ReservationResponse struct {
	HomestayID uint   `json:"homestay_id"`
	Checkin    string `json:"checkin"`
	Checkout   string `json:"checkout"`
}

// var dateLayout1 = "2006-01-02"

func FromCore(dataCore reservation.ReservationCore) ReservationResponse {
	return ReservationResponse{
		HomestayID: dataCore.HomestayID,
		Checkin:    dataCore.Checkin.Format("2006-01-02"),
		Checkout:   dataCore.Checkout.Format("2006-01-02"),
	}
}

func fromCoreAvail(dataCore reservation.ReservationCore) ReservationResponse {
	return ReservationResponse{
		HomestayID: dataCore.ID,
		Checkin:    dataCore.Checkin.Format("2006-01-02"),
		Checkout:   dataCore.Checkout.Format("2006-01-02"),
	}
}
