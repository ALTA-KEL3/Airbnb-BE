package handler

import (
	"airbnb/features/homestay"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type homestayHandler struct {
	srv homestay.HomestayService
}

func New(srv homestay.HomestayService) homestay.HomestayHandler {
	return &homestayHandler{
		srv: srv,
	}
}

// Add implements homestay.HomestayHandler
func (hh *homestayHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := PostHomestayReq{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("image")
		if checkFile != nil {
			formHeader, err := c.FormFile("image")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}

		res, err := hh.srv.Add(token, input.FileHeader, *ReqToCore(input))

		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "file size max 500kb"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "not registered") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success create new homestay",
		})

	}
}

// Delete implements homestay.HomestayHandler
func (*homestayHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// ShowAll implements homestay.HomestayHandler
func (*homestayHandler) ShowAll() echo.HandlerFunc {
	panic("unimplemented")
}

// ShowDetail implements homestay.HomestayHandler
func (*homestayHandler) ShowDetail() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements homestay.HomestayHandler
func (*homestayHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}
