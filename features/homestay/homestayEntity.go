package homestay

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name" validate:"required"`
	Address  string  `json:"address" validate:"required"`
	Phone    string  `json:"phone"`
	Price    float64 `json:"price"`
	Facility string  `json:"facility"`
	Image    string  `json:"image"`
	UserID   uint    `json:"user_id"`
	Feedback Feedback
}

type Feedback struct {
	ID     uint
	Rating int
}

type HomestayHandler interface {
	Add() echo.HandlerFunc
	ShowAll() echo.HandlerFunc
	ShowDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyHomestay() echo.HandlerFunc
}

type HomestayService interface {
	Add(token interface{}, fileData multipart.FileHeader, newHomestay Core) (Core, error)
	ShowAll() ([]Core, error)
	ShowDetail(homestayID uint) (Core, error)
	Update(token interface{}, homestayID uint, fileData multipart.FileHeader, updateData Core) (Core, error)
	Delete(token interface{}, homestayID uint) error
	MyHomestay(token interface{}) ([]Core, error)
}

type HomestayData interface {
	Add(userID uint, newHomestay Core) (Core, error)
	ShowAll() ([]Core, error)
	ShowDetail(homestayID uint) (Core, error)
	Update(userID uint, homestayID uint, updateHomestay Core) (Core, error)
	Delete(userID uint, homestayID uint) error
	MyHomestay(userID uint) ([]Core, error)
}
