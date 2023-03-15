package handler

import (
	"log"
	"net/http"
	"strconv"

	"airbnb/features/feedback"
	"airbnb/helper"

	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	srv feedback.FeedbackServiceInterface
}

func New(srv feedback.FeedbackServiceInterface) feedback.FeedbackHandler {
	return &feedbackHandler{
		srv: srv,
	}
}

func (cc *feedbackHandler) AddFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := FeedbackRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		err = cc.srv.AddFeedback(token, input.HomestayID, *ReqaToCore(input))
		if err != nil {
			log.Println("error running add feedback service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success add feedback",
		})
	}
}

// ListFeedback implements feedback.FeedbackHandler
func (cc *feedbackHandler) ListFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")

		homestayID, _ := strconv.Atoi(paramID)

		res, err := cc.srv.ListFeedback(token, uint(homestayID))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []FeedbackResp{}
		for _, val := range res {
			result = append(result, FeedbackResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success get all user homestays",
		})
	}
}
