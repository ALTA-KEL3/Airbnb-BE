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

import (
	homestay "airbnb/features/homestay/data"
	"airbnb/features/reservation"
	user "airbnb/features/user/data"
	"time"
)

type Reservation struct {
	ID uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	// TenantID          uint
	// OwnerID           uint
	Checkin    time.Time
	Checkout   time.Time
	TotalPrice float64
	HomestayID uint
	UserID     uint
	// ReservationAt     time.Time
	// ReservationStatus string
	TransactionID string
	PaymentURL    string

	// Owner  user.User `gorm:"foreignkey:OwnerID;association_foreignkey:ID"`
	// Tenant user.User `gorm:"foreignkey:BorrowerID;association_foreignkey:ID"`
	User     user.User         `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Homestay homestay.Homestay `gorm:"foreignkey:HomestayID;association_foreignkey:ID"`
}

type HomestayReservation struct {
	ID            uint              `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	ReservationID uint              `json:"reservation_id"`
	HomestayID    uint              `json:"homestay_id"`
	TotalPrice    float64           `json:"total_price"`
	Reservation   Reservation       `gorm:"foreignkey:ReservationID;association_foreignkey:ID"`
	Homestay      homestay.Homestay `gorm:"foreignkey:HomestayID;association_foreignkey:ID"`
}

func ReservationToReservationCore(data Reservation) reservation.ReservationCore {
	return reservation.ReservationCore{
		ID:         data.ID,
		HomestayID: data.HomestayID,
		Checkin:    data.Checkin,
		Checkout:   data.Checkout,
		TotalPrice: data.TotalPrice,
	}
}

func CoreToData(data reservation.ReservationCore) Reservation {
	return Reservation{
		ID:         data.ID,
		HomestayID: data.HomestayID,
		Checkin:    data.Checkin,
		Checkout:   data.Checkout,
		TotalPrice: data.TotalPrice,
	}
}
