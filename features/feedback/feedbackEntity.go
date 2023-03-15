package feedback

import "github.com/labstack/echo/v4"

type FeedbackCore struct {
	ID     uint
	Rating uint
	Note   string
	UserID uint
	// Name       string
	HomestayID    uint
	ReservationID uint
	User          User
}

// type Homestay struct {
// 	ID uint
// 	// Name     string
// 	Feedback []FeedbackCore
// }

type User struct {
	ID             uint
	Name           string
	ProfilePicture string
}

type FeedbackHandler interface {
	AddFeedback() echo.HandlerFunc
	ListFeedback() echo.HandlerFunc
}

type FeedbackServiceInterface interface {
	AddFeedback(token interface{}, homestayID uint, newFeedback FeedbackCore) (FeedbackCore, error)
	ListFeedback(token interface{}, homestayID uint) ([]FeedbackCore, error)
}

type FeedbackDataInterface interface {
	AddFeedback(userID uint, newFeedback FeedbackCore) (FeedbackCore, error)
	ListFeedback(userID uint, homestayID uint) ([]FeedbackCore, error)
}
