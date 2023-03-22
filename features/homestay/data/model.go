package data

import (
	"airbnb/features/homestay"
	"strings"

	// user "airbnb/features/user/data"
	// feedback "airbnb/features/feedback/data"

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
	Feedback []Feedback `gorm:"foreignKey:HomestayRefer"`
}

type Feedback struct {
	gorm.Model
	Rating        uint
	HomestayRefer uint
}

func ModelToCore(data Homestay) homestay.Core {
	return homestay.Core{
		ID:       data.ID,
		Name:     data.Name,
		Address:  data.Address,
		Phone:    data.Phone,
		Price:    data.Price,
		Facility: data.Facility,
		Image:    strings.Split(data.Image, " "),
		UserID:   data.UserID,
		Feedback: homestay.Feedback{},
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
		Image:    strings.Join(data.Image, " "),
		UserID:   data.UserID,
		Feedback: []Feedback{},
	}
}
