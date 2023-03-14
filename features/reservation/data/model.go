package data

// import (
// 	"time"
// 	u "airbnb/features/user/data"
// 	"airbnb/features/reservation"
// 	h "airbnb/features/homestay/data"


// 	"gorm.io/gorm"
// )

// type Reservation struct {
// 	gorm.Model
// 	Checkin    time.Time
// 	Checkout   time.Time
// 	TotalPrice float64
// 	HomestayID uint
// 	Homestay   h.Homestay
// 	UserID     uint
// 	User       u.User
// }

// func CoretoModel(dataCore reservation.ReservationCore) Reservation {
// 	dataGorm := Reservation{
// 		Checkin:    dataCore.Checkin,
// 		Checkout:   dataCore.Checkout,
// 		TotalPrice: dataCore.TotalPrice,
// 		HomestayID: dataCore.HomestayID,
// 		UserID:     dataCore.UserID,
// 	}
// 	return dataGorm
// }
