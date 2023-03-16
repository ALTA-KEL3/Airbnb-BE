package data

import (
	"airbnb/features/homestay"
	user "airbnb/features/user/data"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	Name     string
	Address  string
	Phone    string
	Price    float64
	Facility string
	Image    string
	UserID   uint
	User user.User
}


func ModelToCore(data Homestay) homestay.Core {
	return homestay.Core{
		ID:       data.ID,
		Name:     data.Name,
		Address:  data.Address,
		Phone:    data.Phone,
		Price:    data.Price,
		Facility: data.Facility,
		Image:    data.Image,
		UserID: data.UserID,
	}
}

func CoreToModel(data homestay.Core) Homestay {
	return Homestay{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Address:  data.Address,
		Phone:    data.Phone,
		Price:    data.Price,
		Facility: data.Facility,
		Image:    data.Image,
		UserID:   data.UserID,
	}
}
