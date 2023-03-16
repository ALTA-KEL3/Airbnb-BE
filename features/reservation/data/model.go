package data

import (
	"time"

	"airbnb/features/reservation"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Checkin  time.Time
	Checkout time.Time
	// BookedStart time.Time
	// BookedEnd   time.Time
	Duration   int
	TotalPrice int
	User       User
	Homestay   Homestay
	UserID     uint
	HomestayID uint
	PaymentLink string
}

type User struct {
	gorm.Model
	Name        string
	Reservation []Reservation `gorm:"constraint:OnDelete:CASCADE;"`
}

type Homestay struct {
	gorm.Model
	Name        string
	Address     string
	Price       int
	Reservation []Reservation
}

func FromCore(dataCore reservation.ReservationCore) Reservation {
	reservationGorm := Reservation{
		Checkin:  dataCore.Checkout,
		Checkout: dataCore.Checkout,
		UserID:   dataCore.UserID,
	}
	return reservationGorm
}

func (dataModel *Reservation) toCore() reservation.ReservationCore {
	return reservation.ReservationCore{
		ID:         dataModel.ID,
		Checkin:    dataModel.Checkin,
		Checkout:   dataModel.Checkout,
		TotalPrice: dataModel.Homestay.Price,
		UserID:     dataModel.UserID,
		HomestayID: dataModel.HomestayID,
	}
}

func ToCoreList(dataModel []Reservation) []reservation.ReservationCore {
	var dataCore []reservation.ReservationCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
