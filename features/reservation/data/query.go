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

func (d *reservationQuery) CreateReservation(input reservation.ReservationCore) error {
	model := data.Homestay{}
	tx := d.db.First(&model, input.HomestayID)
	if tx.Error != nil {
		return tx.Error
	}
	// input.Price = model.Price
	input.TotalPrice = r.Checkin(input.Checkin, input.Checkout) * float64(model.Price)

	reservation := CoretoModel(input) //dari gorm model ke user core

	tx1 := d.db.Create(&reservation) // proses insert data

	if tx1.Error != nil {
		return tx.Error
	}
	return nil
}
