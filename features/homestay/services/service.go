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

func (hs *homestayService) Add(token interface{}, fileData multipart.FileHeader, newHomestay homestay.Core) (homestay.Core, error) {

	// if userRole.Role != "hoster"{
	// 	return homestay.Core{}, errors.New("role option only: hoster")
	//

	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return homestay.Core{}, errors.New("user not found")
	}

	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return homestay.Core{}, errors.New("validate: " + err.Error())
	}
	newHomestay.Image1 = url
	newHomestay.Image2 = url
	newHomestay.Image3 = url

	res, err := hs.Data.Add(newHomestay, uint(userID), newHomestay)

	if err != nil {
		log.Println("cannot post book", err.Error())
		return homestay.Core{}, errors.New("server error")
	}

	return res, nil
}

func (hs *homestayService) ShowAll() ([]homestay.Core, error) {
	res, err := hs.Data.ShowAll()

	if err != nil {
		if strings.Contains(err.Error(), "book") {
			return []homestay.Core{}, errors.New("book not found")
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

func (hs *homestayService) Update(token interface{}, homestayID uint, fileData multipart.FileHeader, updatedHomestay homestay.Core) (homestay.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return homestay.Core{}, errors.New("user not found")
	}

	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return homestay.Core{}, errors.New("validate: " + err.Error())
	}
	updatedHomestay.Image1 = url
	updatedHomestay.Image2 = url
	updatedHomestay.Image3 = url

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
