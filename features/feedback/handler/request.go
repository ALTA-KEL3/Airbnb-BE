package handler

import (
	"airbnb/features/feedback"
)

type FeedbackRequest struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	// Name       string `json:"name" form:"name"`
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Rating     uint   `json:"rating" form:"rating"`
	Note       string `json:"note" form:"note"`
}

func RequestToCore(FeedbackInput FeedbackRequest) feedback.FeedbackCore {
	return feedback.FeedbackCore{
		UserID:     FeedbackInput.UserID,
		// Name:       FeedbackInput.Name,
		HomestayID: FeedbackInput.HomestayID,
		Rating:     FeedbackInput.Rating,
		Note:       FeedbackInput.Note,
	}
}


func (r *FeedbackRequest) RequestToCore() *feedback.FeedbackCore {
	return &feedback.FeedbackCore{
		ID:         r.UserID,
		Rating:     r.Rating,
		Note:       r.Note,
		UserID:     r.UserID,
		// Name:       r.Name,
		HomestayID: r.HomestayID,
	}
}

func ReqaToCore(data interface{}) *feedback.FeedbackCore {
	res := feedback.FeedbackCore{}

	switch data.(type) {
	case FeedbackRequest:
		cnv := data.(FeedbackRequest)
		// res.Name = cnv.Name
		res.Rating = cnv.Rating
		res.Note = cnv.Note
		res.UserID = cnv.UserID
		res.HomestayID = cnv.HomestayID

	default:
		return nil
	}
	return &res
}
