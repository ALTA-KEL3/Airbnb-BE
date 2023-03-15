package handler

// package delivery

import (
	// "errors"
	// "log"
	"math"
	"net/http"

	// "strconv"

	"airbnb/features/reservation"
	"airbnb/features/user"
	"airbnb/helper"

	"github.com/labstack/echo/v4"
)

// type ReservationDelivery struct {
// 	reservationService reservation.ReservationServiceInterface
// }

// func New(service reservation.ReservationServiceInterface, e *echo.Echo) {
// 	handler := &ReservationDelivery{
// 		reservationService: service,
// 	}

type reservationHandler struct {
	srv reservation.ReservationServiceInterface
	// rsvSrv user.UserService
}

func New(srv reservation.ReservationServiceInterface, rsvSrvc user.UserService) reservation.ReservationHandler {
	return &reservationHandler{
		srv: srv,
		// rsvSrv: rsvSrvc,
	}
}

// e.POST("/reservations/check", handler.CheckAvailability, middlewares.JWTMiddleware())
// e.POST("/reservations", handler.Payment, middlewares.JWTMiddleware())
// e.GET("/reservations", handler.GetHistory, middlewares.JWTMiddleware())
// }
// func (hh *homestayHandler) ShowAll() echo.HandlerFunc {
func (d *reservationHandler) CheckAvailability() echo.HandlerFunc {

	return func(c echo.Context) error {

		input := ReservationRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusNotFound, helper.ResponseFail("requested resource was not found"+errBind.Error()))
		}
		dataInput := ToCore(input)
		res, err := d.srv.CheckAvailability(dataInput)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFail("error read data"))
		}

		start := dataInput.Checkin
		end := dataInput.Checkout
		period := int(math.Ceil(end.Sub(start).Hours() / 24))

		dataResponse := fromCoreAvail(res)
		dataResponse.Duration = period
		dataResponse.TotalPrice = dataInput.Duration * dataResponse.Price

		return c.JSON(http.StatusOK, helper.ResponseSuccess("available reservation", dataResponse))

	}
}

// func (d *ReservationDelivery) Payment(c echo.Context) error {
// 	idUser := middlewares.ExtractTokenUserId(c)
// 	input := PaymentRequest{}
// 	errBind := c.Bind(&input)
// 	if errBind != nil {
// 		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
// 	}

// 	input.UserID = uint(idUser)
// 	dataInput := ToCorePayment(input)
// 	err := d.reservationService.CreatePayment(dataInput)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
// 	}
// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Success reservation, see you later"))
// }

// func (d *ReservationDelivery) GetHistory(c echo.Context) error {
// 	idUser := middlewares.ExtractTokenUserId(c)
// 	userId := uint(idUser)
// 	results, err := d.reservationService.GetHistory(userId)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
// 	}

// 	dataResponse := TripArr(results)

// 	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get trip history.", dataResponse))
// }
