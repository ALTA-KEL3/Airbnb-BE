package handler

// import (
// 	"airbnb/features/reservation"
// 	"time"
// )

// type ReservationRequest struct {
// 	Checkin    string  `json:"checkin" form:"checkin"`
// 	Checkout   string  `json:"checkout" form:"checkout"`
// 	TotalPrice float64 `json:"total_price " form:"total_price"`
// 	HomestayID uint    `json:"homestay_id" form:"homestay_id"`
// 	UserID     uint    `json:"user_id" form:"user_id"`
// }

// func RequestToCore(data ReservationRequest, Checkin, Checkout time.Time) reservation.ReservationCore {
// 	return reservation.ReservationCore{
// 		HomestayID: data.HomestayID,
// 		TotalPrice: data.TotalPrice,
// 		UserID:     data.UserID,
// 		Checkin:    Checkin,
// 		Checkout:   Checkout,
// 	}
// }


import (
	// "math"
	"time"

	"airbnb/features/reservation"
)

type ReservationRequest struct {
	Checkin    string `json:"checkin" form:"checkin"`
	Checkout   string `json:"checkout" form:"checkout"`
	HomestayID int    `json:"homestay_id" form:"homestay_id"`
	Homestay   Homestay
}

type Homestay struct {
	ID          uint
	BookedStart time.Time
	BookedEnd   time.Time
}


var dateLayout = "2006-01-02"


func ToCore(reservationInput ReservationRequest) reservation.ReservationCore {
	in, _ := time.Parse(dateLayout, reservationInput.Checkin)
	out, _ := time.Parse(dateLayout, reservationInput.Checkout)
	return reservation.ReservationCore{
		Checkin:    in,
		Checkout:   out,
		HomestayID: uint(reservationInput.HomestayID),
		Homestay: reservation.Homestay{
			ID:          reservationInput.Homestay.ID,
			// BookedStart: reservationInput.Homestay.BookedStart,
			BookedEnd:   reservationInput.Homestay.BookedEnd,
		},
	}
}

