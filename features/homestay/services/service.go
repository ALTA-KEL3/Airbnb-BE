package services

import (
	"airbnb/features/homestay"
	"airbnb/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
)

type homestayService struct {
	Data homestay.HomestayData
	// Data     homestay.HomestayService
	// validate *validator.Validate
}

func New(data homestay.HomestayData) homestay.HomestayService {
	return &homestayService{
		Data: data,
		// validate: validator.New(),
	}
}

func (hs *homestayService) Add(token interface{}, images []*multipart.FileHeader, newHomestay homestay.Core) (homestay.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return homestay.Core{}, errors.New("user not found")
	}

	var urls []string
	for i, image := range images {
		url, err := helper.GetUrlImagesFromAWS(*image, i)
		if err != nil {
			return homestay.Core{}, errors.New("validate: " + err.Error())
		}
		urls = append(urls, url)
	}

	newHomestay.Image = urls

	res, err := hs.Data.Add(uint(userID), newHomestay)

	if err != nil {
		log.Println("cannot post homestay", err.Error())
		return homestay.Core{}, errors.New("server error")
	}

	return res, nil
}

func (hs *homestayService) ShowAll() ([]homestay.Core, error) {
	res, err := hs.Data.ShowAll()

	if err != nil {
		if strings.Contains(err.Error(), "homestay") {
			return []homestay.Core{}, errors.New("homestay not found")
		} else {
			return []homestay.Core{}, errors.New("internal server error")
		}
	}
	return res, nil
}

func (hs *homestayService) ShowDetail(homestayID uint) (homestay.Core, error) {
	res, err := hs.Data.ShowDetail(homestayID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return homestay.Core{}, errors.New("data not found")
		} else {
			return homestay.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (hs *homestayService) Update(token interface{}, homestayID uint, images []*multipart.FileHeader, updatedHomestay homestay.Core) (homestay.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return homestay.Core{}, errors.New("user not found")
	}

	var urls []string
	for i, image := range images {
		url, err := helper.GetUrlImagesFromAWS(*image, i)
		if err != nil {
			return homestay.Core{}, errors.New("validate: " + err.Error())
		}
		urls = append(urls, url)
	}
	updatedHomestay.Image = urls

	res, err := hs.Data.Update(uint(userID), homestayID, updatedHomestay)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return homestay.Core{}, errors.New("data not found")
		} else {
			return homestay.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (hs *homestayService) Delete(token interface{}, homestayID uint) error {
	id := helper.ExtractToken(token)

	err := hs.Data.Delete(uint(id), homestayID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("data not found")
	}

	return nil
}

func (hs *homestayService) MyHomestay(token interface{}) ([]homestay.Core, error) {
	id := helper.ExtractToken(token)

	res, err := hs.Data.MyHomestay(uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "homestay not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []homestay.Core{}, errors.New(msg)
	}

	return res, nil
}
