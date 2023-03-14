package data

import (
	"errors"

	"airbnb/features/feedback"
	"log"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackDataInterface {
	return &feedbackQuery{
		db: db,
	}
}

func (cq *feedbackQuery) AddFeedback(userID uint, newFeedback feedback.FeedbackCore) (feedback.FeedbackCore, error) {
	cnv := FeedbackCoreToFeedback(newFeedback)
	cnv.UserID = uint(userID)

	err := cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return feedback.FeedbackCore{}, errors.New("server error")
	}
	newFeedback.ID = cnv.ID
	newFeedback.HomestayID = cnv.HomestayID

	return newFeedback, nil
}
