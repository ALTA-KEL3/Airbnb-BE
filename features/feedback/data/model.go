package data

import (
	"airbnb/features/feedback"

	// homestay "airbnb/features/homestay/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Rating uint
	Note   string
	UserID uint
	User   User
	// Name       string
	HomestayID uint
	// Homestay   homestay.Homestay
	// Homestay Homestay
	ReservationID uint
}

type User struct {
	gorm.Model
	Name           string
	ProfilePicture string
	// Feedback []Feedback `gorm:"foreignKey:UserID"`
}
type Homestay struct {
	gorm.Model
	// ID uint
	Name     string
	UserID   uint
	Rating   uint
	Feedback string
	// Feedback []Feedback `gorm:"foreignKey:HomestayID"`
}

func FeedbackDataToCore(data Feedback) feedback.FeedbackCore {
	return feedback.FeedbackCore{
		ID:     data.ID,
		Rating: data.Rating,
		Note:   data.Note,
		UserID: data.UserID,
		// Name:       data.Name,
		HomestayID:    data.HomestayID,
		ReservationID: data.ReservationID,
		User: feedback.User{
			ID:             data.User.ID,
			Name:           data.User.Name,
			ProfilePicture: data.User.ProfilePicture,
		},
	}

}

func FeedbackCoreToFeedback(dataCore feedback.FeedbackCore) Feedback {
	return Feedback{
		Model:  gorm.Model{ID: dataCore.ID},
		Rating: dataCore.Rating,
		Note:   dataCore.Note,
		UserID: dataCore.UserID,
		// Name:       dataCore.Name,
		HomestayID:    dataCore.HomestayID,
		ReservationID: dataCore.ReservationID,
		User: User{
			Name:           dataCore.User.Name,
			ProfilePicture: dataCore.User.ProfilePicture,
		},
	}
}

// func FeedbackDataToCore(data Feedback) feedback.FeedbackCore {
// 	return feedback.FeedbackCore{
// 		ID:            data.ID,
// 		Rating:        data.Rating,
// 		Note:          data.Note,
// 		UserID:        data.UserID,
// 		HomestayID:    data.HomestayID,
// 		ReservationID: data.ReservationID,
// 		User:          feedback.User{},
// 	}
// }

// func FeedbackCoreToFeedback(data feedback.FeedbackCore) Feedback {
// 	return Feedback{
// 		Model:         gorm.Model{ID: data.ID},
// 		Rating:        data.Rating,
// 		Note:          data.Note,
// 		UserID:        data.UserID,
// 		HomestayID:    data.HomestayID,
// 		ReservationID: data.ReservationID,
// 		User:          User{},
// 	}
// }
