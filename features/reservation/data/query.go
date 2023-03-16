package data

import (
	"airbnb/features/reservation"

	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationDataInterface {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) CheckAvailability(input reservation.ReservationCore) (data reservation.ReservationCore , err error) {

	var reservation Reservation

	tx := r.db.Where("homestay_id=? AND ? BETWEEN checkin AND checkout OR ? BETWEEN checkout AND checkout", input.HomestayID, input.Checkin, input.Checkout).First(&reservation, input.HomestayID) //

	if tx.Error != nil {
		return data, tx.Error
	}
	data = reservation.toCore()
	return data, nil
}

// SELECT * FROM reservations WHERE room_id = ? AND ? BETWEEN booked_start AND booked_end OR ? BETWEEN booked_start AND booked_end
// Select(“users.name, profiles.bio”).Joins(“INNER JOIN profiles ON profiles.user_id = users.id”).Find(&users)
// db.Table("table1").Joins("INNER JOIN table2 ON table1.col2 = table2.col2").Select("table1.col1, table1.col2, table2.col3").Where("table1.col1 = ?", input).Find(&results).
