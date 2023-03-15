package handler

// package delivery

import (
	"time"

	"airbnb/features/reservation"
)

type ReservationResponse struct {
	ID         uint `json:"id"`
	Duration   int  `json:"duration"`
	TotalPrice int  `json:"total_price"`
}

type HomestayResponse struct {
	HomestayID uint `json:"homestay_id"`
	Price      int  `json:"price"`
	Duration   int  `json:"duration"`
	TotalPrice int  `json:"total_price"`
}

type HomestayTrip struct {
	Name    string `json:"homestay_name"`
	Address string `json:"address"`
}
type HistoryResponse struct {
	Checkin         time.Time `json:"checkin"`
	Checkout        time.Time `json:"checkout"`
	Duration        int       `json:"duration"`
	TotalPrice      int       `json:"total_price"`
	HomestayName    string    `json:"homestay_name"`
	HomestayAddress string    `json:"homestay_address"`
	HomestayID      uint      `json:"homestay_id"`
}

func FromCore(dataCore reservation.ReservationCore) ReservationResponse {
	return ReservationResponse{
		ID:         dataCore.ID,
		Duration:   dataCore.Duration,
		TotalPrice: dataCore.TotalPrice,
	}
}

func fromCoreAvail(dataCore reservation.Homestay) HomestayResponse {
	return HomestayResponse{
		HomestayID: dataCore.ID,
		Price:      dataCore.Price,
	}
}

// func fromCoreTrip(dataCore reservation.ReservationCore) HistoryResponse {
// 	return HistoryResponse{
// 		StartDate:       dataCore.Checkin,
// 		EndDate:         dataCore.Checkout,
// 		Duration:        dataCore.Duration,
// 		TotalPrice:      dataCore.TotalPrice,
// 		HomestayID:      dataCore.HomestayID,
// 		HomestayName:    dataCore.Homestay.Name,
// 		HomestayAddress: dataCore.Homestay.Address,
// 	}
// }

// func TripArr(dataCore []reservation.ReservationCore) []HistoryResponse {
// 	var dataResponse []HistoryResponse
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCoreTrip(v))
// 	}
// 	return dataResponse
// }
