package handler

import (
	"airbnb/features/homestay"
	"airbnb/features/user"
	"airbnb/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type homestayHandler struct {
	srv    homestay.HomestayService
	usrSrv user.UserService
}

func New(srv homestay.HomestayService, usrSrvc user.UserService) homestay.HomestayHandler {
	return &homestayHandler{
		srv:    srv,
		usrSrv: usrSrvc,
	}
}

// Add implements homestay.HomestayHandler
func (hh *homestayHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := PostHomestayReq{}

		user, errr := hh.usrSrv.Profile(token)
		if errr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})

		}
		if user.Role != "hoster" {
			log.Println("access denied, cannot add product because you are not hoster")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "access denied"})
		}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile1, _, _ := c.Request().FormFile("image1")
		if checkFile1 != nil {
			formHeader1, err := c.FormFile("image1")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader1 = *formHeader1
		}

		checkFile2, _, _ := c.Request().FormFile("image2")
		if checkFile2 != nil {
			formHeader2, err := c.FormFile("image2")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader2 = *formHeader2
		}

		checkFile3, _, _ := c.Request().FormFile("image3")
		if checkFile3 != nil {
			formHeader3, err := c.FormFile("image3")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader3 = *formHeader3
		}

		res, err := hh.srv.Add(token, input.FileHeader1, input.FileHeader2, input.FileHeader3, *ReqToCore(input))

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

		res.UserID = user.ID

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success create new homestay",
		})

	}
}

// Delete implements homestay.HomestayHandler
func (hh *homestayHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("homestay_id")

		homestayID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		err = hh.srv.Delete(token, uint(homestayID))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete homestay",
		})
	}
}

// ShowAll implements homestay.HomestayHandler
func (hh *homestayHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := hh.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		result := []ShowAllHomestay{}
		for _, val := range res {
			result = append(result, ShowAllHomestayJson(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all homestay",
		})
	}
}

// ShowDetail implements homestay.HomestayHandler
func (hh *homestayHandler) ShowDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("homestay_id")

		homestayID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		res, err := hh.srv.ShowDetail(uint(homestayID))

		if err != nil {
			if strings.Contains(err.Error(), "data") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "data not found",
				})
			}
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "success get homestay details", HomestayResponse(res)))
	}
}

// Update implements homestay.HomestayHandler
func (hh *homestayHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("homestay_id")

		homestayID, err := strconv.Atoi(paramID)

		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "Invalid input")
		}

		input := PostHomestayReq{}

		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		//proses cek apakah user input foto ?
		checkFile, _, _ := c.Request().FormFile("image1")
		if checkFile != nil {
			formHeader1, err := c.FormFile("image1")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader1 = *formHeader1

			formHeader2, err := c.FormFile("image2")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader2 = *formHeader2

			formHeader3, err := c.FormFile("image3")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader3 = *formHeader3
		}

		res, err := hh.srv.Update(token, uint(homestayID), input.FileHeader1, input.FileHeader2, input.FileHeader3, *ReqToCore(input))

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

		result, err := ConvertHomestayUpdateResponse(res)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			// log.Println(res)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    result,
				"message": "success update homestay",
			})
		}
	}
}

func (hh *homestayHandler) MyHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		res, err := hh.srv.MyHomestay(token)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []Homestay{}
		for _, val := range res {
			result = append(result, HomestayResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success get all user homestays",
		})
	}
}
