package data

import (
	"airbnb/features/homestay/data"
	"airbnb/features/reservation"
	r "airbnb/features/reservation/services"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

func NewReservation(db *gorm.DB) reservation.ReservationDataInterface {
	return &reservationQuery{
		db: db,
	}
}

func (d *reservationQuery) CreateReservation(input reservation.ReservationCore) ( reservation.ReservationCore, error) {
	model := data.Homestay{}
	tx := d.db.First(&model, input.HomestayID)
	if tx.Error != nil {
		return reservation.ReservationCore{}, tx.Error
	}
	// input.Price = model.Price
	input.TotalPrice = r.Checkin(input.Checkin, input.Checkout) * float64(model.Price)

	res := CoretoModel(input) //dari gorm model ke user core

	err := d.db.Create(&res).Error // proses insert data

	if err != nil {
		return reservation.ReservationCore{}, tx.Error
	}
	return input, nil
}
