package handler

import (
	"log"
	"net/http"

	"airbnb/features/feedback"

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

		res, err := cc.srv.AddFeedback(token, input.HomestayID, *ReqaToCore(input))
		if err != nil {
			log.Println("error running add feedback service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success add book to feedback",
		})
	}
}

