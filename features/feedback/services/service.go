package services

import (
	"airbnb/features/feedback"
	"airbnb/helper"
	"errors"
	"log"

	// "mime/multipart"
	"strings"
)

type feedbackService struct {
	qry feedback.FeedbackDataInterface
}

func New(cd feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &feedbackService{
		qry: cd,
	}
}

func (cuc *feedbackService) AddFeedback(token interface{}, homestayID uint, newFeedback feedback.FeedbackCore) (feedback.FeedbackCore, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token add feedback")
		return feedback.FeedbackCore{}, errors.New("user not found")
	}
	res, err := cuc.qry.AddFeedback(uint(userId), newFeedback)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return feedback.FeedbackCore{}, errors.New(msg)
	}
	return res, nil
}
