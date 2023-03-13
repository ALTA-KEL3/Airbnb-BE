package service

import (
	"airbnb/features/reservation"
	"errors"
	"time"
	// "github.com/go-playground/validator/v10"
)

type ReservationService struct {
	Data reservation.ReservationDataInterface
	// validasi *validator.Validate
}

func NewReservation(data reservation.ReservationDataInterface) reservation.ReservationServiceInterface {
	return &ReservationService{
		Data: data,
		// validate: validator.New(),
	}

}

func (service *ReservationService) CreateReservation(input reservation.ReservationCore) error {
	err := service.Data.CreateReservation(input)

	if err != nil {
		return errors.New("failed to add reservation")
	}
	return nil
}
func Checkin(start, end time.Time) float64 {
	difference := start.Sub(end)
	days := float64(difference.Hours() / 24)

	return days

}
