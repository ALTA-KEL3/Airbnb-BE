package service

import (
	"airbnb/features/reservation"
// 	"airbnb/helper"
	"errors"
	// "time"
// 	// "github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservation.ReservationDataInterface
}

func New(data reservation.ReservationDataInterface) reservation.ReservationServiceInterface {
	return &reservationService{
		reservationData: data,
	}
}

// func (s *reservationService) CheckAvailability(input reservation.ReservationCore) (data reservation.Homestay, err error) {
// 	data, err = s.reservationRepo.CheckAvailability(input)
// 	if err != nil {
// 		return data, errors.New("failed get data, error query")

// 	}
// 	return data, nil
// }

func (s *reservationService) CheckAvailability(input reservation.ReservationCore) (data reservation.Homestay,err error) {
	
	data, err = s.reservationData.CheckAvailability(input)
	if err != nil {
		return data, errors.New("failed get data, error query")

	}
	return data, nil
}