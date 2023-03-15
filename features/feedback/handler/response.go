package handler

import (
	"airbnb/features/feedback"
)

type FeedbackResp struct {
	ID         uint   `json:"id"`
	Rating     uint   `json:"rating"`
	Note       string `json:"note"`
	UserID     uint   `json:"user_id"`
	HomestayID uint   `json:"homestay_id"`
}

func FeedbackResponse(data feedback.FeedbackCore) FeedbackResp {
	return FeedbackResp{
		ID:         data.ID,
		Rating:     data.Rating,
		Note:       data.Note,
		UserID:     data.UserID,
		HomestayID: data.HomestayID,
	}
}
