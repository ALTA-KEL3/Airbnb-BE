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

func (cuc *feedbackService) AddFeedback(token interface{}, homestayID uint, newFeedback feedback.FeedbackCore) error {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("error extract token add feedback")
		return errors.New("user not found")
	}
	err := cuc.qry.AddFeedback(uint(userId), newFeedback)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
}

// ListFeedback implements feedback.FeedbackServiceInterface
func (cuc *feedbackService) ListFeedback(token interface{}, homestayID uint) ([]feedback.FeedbackCore, error) {
	id := helper.ExtractToken(token)

	res, err := cuc.qry.ListFeedback(uint(id), uint(homestayID))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "feedback not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []feedback.FeedbackCore{}, errors.New(msg)
	}

	return res, nil
}
