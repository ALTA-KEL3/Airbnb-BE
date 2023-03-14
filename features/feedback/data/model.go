package data

import (
	"airbnb/features/feedback"
	// homestay "airbnb/features/homestay/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Rating     uint
	Note       string
	UserID     uint
	User       User
	// Name       string
	HomestayID uint
	// Homestay   homestay.Homestay
	Homestay Homestay
}

type User struct {
	gorm.Model
	Feedback []Feedback `gorm:"foreignKey:UserID"`
}
type Homestay struct {
	gorm.Model
	// ID uint
	// Name     string
	Feedback []Feedback `gorm:"foreignKey:HomestayID"`
}

func FeedbackCoreToFeedback(dataCore feedback.FeedbackCore) Feedback {
	feedbackGorm := Feedback{
		Model:      gorm.Model{ID: dataCore.ID},
		Rating:     dataCore.Rating,
		Note:       dataCore.Note,
		UserID:     dataCore.UserID,
		// Name:       dataCore.Name,
		HomestayID: dataCore.HomestayID,
	}
	return feedbackGorm
}

func FeedbackDataToCore(data Feedback) feedback.FeedbackCore {
	dataFeedback := feedback.FeedbackCore{
		ID:         data.ID,
		Rating:     data.Rating,
		Note:       data.Note,
		UserID:     data.UserID,
		// Name:       data.Name,
		HomestayID: data.HomestayID,
	}
	return dataFeedback
}
