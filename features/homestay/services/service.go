package services

import (
	"airbnb/features/homestay"
	"airbnb/features/user"
	"airbnb/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"
)

type homestayService struct {
	// qry homestay.HomestayData
	Data     homestay.HomestayService
	// validate *validator.Validate
}

func New(data homestay.HomestayData) homestay.HomestayService {
	return &homestayService{
		Data:     data,
		// validate: validator.New(),
	}
}

func (hs *homestayService) Add(userRole user.Core, token interface{}, fileData multipart.FileHeader, newHomestay homestay.Core) (homestay.Core, error) {
	
	if userRole.Role != "hoster"{
		return homestay.Core{}, errors.New("role option only: hoster")
	}


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

func (hs *homestayService) Update(token interface{}, homestayID uint, fileData multipart.FileHeader, updatedBook homestay.Core) (homestay.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return book.Core{}, errors.New("user not found")
	}

	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return book.Core{}, errors.New("validate: " + err.Error())
	}
	updatedBook.Image = url
	res, err := buc.qry.Update(uint(userID), bookID, updatedBook)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return book.Core{}, errors.New("data not found")
		} else {
			return book.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (buc *bookUseCase) Delete(token interface{}, bookID uint) error {
	id := helper.ExtractToken(token)

	err := buc.qry.Delete(uint(id), bookID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("data not found")
	}

	return nil
}

func (buc *bookUseCase) MyBook(token interface{}) ([]book.Core, error) {
	id := helper.ExtractToken(token)
	res, err := buc.qry.MyBook(uint(id))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []book.Core{}, errors.New(msg)
	}

	return res, nil
}
