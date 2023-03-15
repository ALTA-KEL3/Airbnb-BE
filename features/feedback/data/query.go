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

// ListFeedback implements feedback.FeedbackDataInterface
func (cq *feedbackQuery) ListFeedback(userID uint, homestayID uint) ([]feedback.FeedbackCore, error) {
	res := []Feedback{}
	if err := cq.db.Where("user_id = ?", userID).Order("created_at desc").Find(&res).Error; err != nil {
		log.Println("get feedback data query error : ", err.Error())
		return []feedback.FeedbackCore{}, err
	}
	result := []feedback.FeedbackCore{}
	for _, val := range res {
		result = append(result, FeedbackDataToCore(val))
	}

	return result, nil
}
