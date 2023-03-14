package handler

import (
	"airbnb/features/homestay"
	"mime/multipart"
)

type PostHomestayReq struct {
	Name       string  `json:"name" form:"name"`
	Address    string  `json:"address" form:"address"`
	Phone      string  `json:"phone" form:"phone"`
	Price      float64 `json:"price" form:"price"`
	Facility   string  `json:"facility" form:"facility"`
	FileHeader multipart.FileHeader
}

// type UserCore struct {
// 	ID             uint   `json:"id" form:"id"`
// 	Name           string `json:"name" form:"name"`
// 	Address        string `json:"address" form:"address"`
// 	Phone          string `json:"phone" form:"phone"`
// 	Role           string `json:"role" form:"role"`
// 	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
// }

func ReqToCore(data interface{}) *homestay.Core {
	res := homestay.Core{}

	switch data.(type) {
	case PostHomestayReq:
		cnv := data.(PostHomestayReq)
		res.Name = cnv.Name
		res.Address = cnv.Address
		res.Phone = cnv.Phone
		res.Price = cnv.Price
		res.Facility = cnv.Facility
	default:
		return nil
	}
	return &res
}
