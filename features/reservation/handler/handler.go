package handler

import (

	"net/http"

	"airbnb/features/reservation"
	"airbnb/features/user"
	"airbnb/helper"

	"github.com/labstack/echo/v4"
)

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

		start := dataInput.Checkin.Format("2006-01-02")
		end := dataInput.Checkout.Format("2006-01-02")

		dataResponse := fromCoreAvail(res)
		dataResponse.HomestayID = dataInput.HomestayID
		dataResponse.Checkin = start
		dataResponse.Checkout = end

		return c.JSON(http.StatusOK, helper.ResponseSuccess("available reservation", dataResponse))

	}
}
