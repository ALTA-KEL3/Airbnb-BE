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
	// Role     string
	User user.User
}

// type UserCore struct {
// 	gorm.Model
// 	Name           string
// 	Address        string
// 	Phone          string
// 	Role           string
// 	ProfilePicture string
// }

func ModelToCore(data Homestay) homestay.Core {
	return homestay.Core{
		ID:       data.ID,
		Name:     data.Name,
		Address:  data.Address,
		Phone:    data.Phone,
		Price:    data.Price,
		Facility: data.Facility,
		Image:    data.Image,
		// Role:     data.Role,
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
		// Role:     data.Role,
	}
}
