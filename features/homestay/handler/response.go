package handler

import (
	"airbnb/features/homestay"
	"errors"
)

type Homestay struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Price    float64  `json:"price"`
	Facility string   `json:"facility"`
	Image    []string `json:"image"`
	Rating   uint     `json:"rating"`
}

func HomestayResponse(data homestay.Core) Homestay {
	return Homestay{
		ID:       data.ID,
		Name:     data.Name,
		Address:  data.Address,
		Phone:    data.Phone,
		Price:    data.Price,
		Facility: data.Facility,
		Image:    data.Image,
		Rating:   data.Feedback.Rating,
	}
}

type ShowAllHomestay struct {
	ID       uint     `json:"id"`
	Image    []string `json:"image"`
	Name     string   `json:"name"`
	Facility string   `json:"facility"`
	Price    float64  `json:"price"`
	Rating   uint     `json:"rating"`
}

func ShowAllHomestayJson(data homestay.Core) ShowAllHomestay {
	return ShowAllHomestay{
		ID:       data.ID,
		Image:    data.Image,
		Name:     data.Name,
		Facility: data.Facility,
		Price:    data.Price,
		Rating:   data.Feedback.Rating,
	}
}

func ConvertHomestayUpdateResponse(input homestay.Core) (interface{}, error) {
	ResponseFilter := homestay.Core{}
	ResponseFilter = input
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.Name != "" {
		result["name"] = ResponseFilter.Name
	}
	if ResponseFilter.Address != "" {
		result["address"] = ResponseFilter.Address
	}
	if ResponseFilter.Phone != "" {
		result["phone"] = ResponseFilter.Phone
	}
	if ResponseFilter.Price != 0 {
		result["price"] = ResponseFilter.Price
	}
	if ResponseFilter.Facility != "" {
		result["facility"] = ResponseFilter.Facility
	}
	if len(ResponseFilter.Image) != 0 {
		result["image"] = ResponseFilter.Image
	}
	if len(result) < 1 {
		return homestay.Core{}, errors.New("no data was change")
	}
	return result, nil
}
