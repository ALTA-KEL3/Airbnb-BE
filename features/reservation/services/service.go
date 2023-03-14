package service

// import (
// 	"airbnb/features/reservation"
// 	"airbnb/helper"
// 	"errors"
// 	"time"
// 	// "github.com/go-playground/validator/v10"
// )

// type ReservationService struct {
// 	Data reservation.ReservationDataInterface
// 	// validasi *validator.Validate
// }

// func NewReservation(data reservation.ReservationDataInterface) reservation.ReservationServiceInterface {
// 	return &ReservationService{
// 		Data: data,
// 		// validate: validator.New(),
// 	}

// }

// func (service *ReservationService) CreateReservation(token interface{}, totalPrice float64) (reservation.ReservationCore, string, error) {
// 	// var token interface{}
// 	userID := helper.ExtractToken(token)

// 	if userID <= 0 {
// 		return reservation.ReservationCore{}, errors.New("user not found")
// 	}
// 	res, err := service.Data.CreateReservation(input)
// 	if err != nil {
// 		return reservation.ReservationCore{}, errors.New("failed to add reservation")
// 	}
// 	return res, nil
// }
// func Checkin(start, end time.Time) float64 {
// 	difference := start.Sub(end)
// 	days := float64(difference.Hours() / 24)

// 	return days

// }
