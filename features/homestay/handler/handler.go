package handler

import (
	"airbnb/features/feedback"
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
	fdbSrv feedback.FeedbackServiceInterface
}

func New(srv homestay.HomestayService, usrSrvc user.UserService, fdbSrvc feedback.FeedbackServiceInterface) homestay.HomestayHandler {
	return &homestayHandler{
		srv:    srv,
		usrSrv: usrSrvc,
		fdbSrv: fdbSrvc,
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
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "access deniedcannot add product because you are not hoster"})
		}

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
		token := c.Get("user")
		rhomestay, err := hh.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		result := []ShowAllHomestay{}
		for _, val := range rhomestay {
			rating, err := hh.getRating(token, val.ID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
			val.Feedback.Rating = uint(rating)
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
		token := c.Get("user")

		homestayID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error", err.Error())
			return c.JSON(http.StatusBadGateway, "invalid input")
		}

		homestay, _ := hh.srv.ShowDetail(uint(homestayID))

		rating, err := hh.getRating(token, uint(homestayID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		homestay.Feedback.Rating = uint(rating)

		if err != nil {
			if strings.Contains(err.Error(), "data") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "data not found",
				})
			}
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "success get homestay details", HomestayResponse(homestay)))
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
		checkFile, _, _ := c.Request().FormFile("image")
		if checkFile != nil {
			formHeader, err := c.FormFile("image")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}

		res, err := hh.srv.Update(token, uint(homestayID), input.FileHeader, *ReqToCore(input))

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
			rating, err := hh.getRating(token, val.ID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
			val.Feedback.Rating = uint(rating)
			result = append(result, HomestayResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success get all user homestays",
		})
	}
}

func (hh *homestayHandler) getRating(token interface{}, homestayID uint) (uint, error) {
	sum := uint(0)
	rFeedback, err := hh.fdbSrv.ListFeedback(token, homestayID)
	if err != nil {
		return 0, err
	}
	for _, feedback := range rFeedback {
		sum += uint(feedback.Rating)
		if len(rFeedback) == 0 {
			log.Println("No ratings found.")
			return 0, nil
		}
	}

	if sum == 0 {
		return 0, nil
	}
	return uint(sum) / uint(len(rFeedback)), nil
}

// sudo
