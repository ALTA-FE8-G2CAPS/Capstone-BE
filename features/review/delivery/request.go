package delivery

import "capstone-project/features/review"

type ReviewRequest struct {
	VenueID     uint   `json:"venue_id" form:"venue_id"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Rate        uint   `json:"rate" form:"rate"`
	Feedback    string `json:"feedback" form:"feedback"`
	Foto_review string `json:"foto_review" form:"foto_review"`
}

func ToCore(data ReviewRequest) review.ReviewCore {
	return review.ReviewCore{
		VenueID:     data.VenueID,
		UserID:      data.UserID,
		Rate:        data.Rate,
		Feedback:    data.Feedback,
		Foto_review: data.Foto_review,
	}
}
