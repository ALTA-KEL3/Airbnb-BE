package feedback

import "github.com/labstack/echo/v4"

type FeedbackCore struct {
	ID         uint
	Rating     uint
	Note       string
	UserID     uint
	// Name       string
	HomestayID uint
}

type Homestay struct {
	ID       uint
	// Name     string
	Feedback []FeedbackCore
}

type FeedbackHandler interface {
	AddFeedback() echo.HandlerFunc
}

type FeedbackServiceInterface interface {
	AddFeedback(token interface{}, homestayID uint, newFeedback FeedbackCore) (FeedbackCore, error)
}

type FeedbackDataInterface interface {
	AddFeedback(userID uint, newFeedback FeedbackCore) (FeedbackCore, error)
}

