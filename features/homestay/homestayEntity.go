package homestay

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Name     string `validate:"required"`
	Address  string `validate:"required"`
	Phone    string `validate:"required"`
	Price    string
	Facility string
	Image1   string
	Image2   string
	Image3   string
	UserID   uint
}

type HomestayHandler interface {
	Add() echo.HandlerFunc
	ShowAll() echo.HandlerFunc
	ShowDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type HomestayService interface {
	Add(token interface{}, fileData multipart.FileHeader, newHomestay Core) (Core, error)
	ShowAll() ([]Core, error)
	ShowDetail(homestayID uint) (Core, error)
	Update(token interface{}, homestayID uint, fileData multipart.FileHeader, updateData Core) (Core, error)
	Delete(token interface{}, homestayID uint) error
}

type HomestayData interface {
	Add(userRole Core, userID uint, newHomestay Core) (Core, error)
	ShowAll() ([]Core, error)
	ShowDetail(homestayID uint) (Core, error)
	Update(userID uint, homestayID uint, updateHomestay Core) (Core, error)
	Delete(userID uint, homestayID uint) error
}