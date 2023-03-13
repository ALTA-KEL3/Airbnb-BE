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
	Phone    string `gorm:"unique"`
	Price    float64
	Facility string
	Image1   string
	Image2   string
	Image3   string
	UserID   uint
	User     user.User
}

func ModelToCore(data Homestay) homestay.Core {
	return homestay.Core{
		ID:       data.ID,
		Name:     data.Name,
		Address:  data.Address,
		Phone:    data.Phone,
		Price:    data.Price,
		Facility: data.Facility,
		Image1:   data.Image1,
		Image2:   data.Image2,
		Image3:   data.Image3,
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
		Image1:   data.Image1,
		Image2:   data.Image2,
		Image3:   data.Image3,
		UserID:   data.UserID,
	}
}
