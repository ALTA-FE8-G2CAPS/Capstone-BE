package delivery

import "capstone-project/features/review"

type ReviewResponse struct {
	ID          uint   `json:"id"`
	Name_user   string `json:"name_user"`
	Foto_user   string `json:"foto_user"`
	VenueID     uint   `json:"venue_id"`
	Nama_venue  string `json:"nama_venue"`
	UserID      uint   `json:"user_id"`
	Rate        uint   `json:"rate"`
	Feedback    string `json:"feedback"`
	Foto_review string `json:"foto_review"`
}

func FromCore(data review.ReviewCore) ReviewResponse {
	return ReviewResponse{
		ID:          data.ID,
		Name_user:   data.Name_user,
		Foto_user:   data.Foto_user,
		VenueID:     data.VenueID,
		Nama_venue:  data.Nama_venue,
		UserID:      data.UserID,
		Rate:        data.Rate,
		Feedback:    data.Feedback,
		Foto_review: data.Foto_review,
	}
}

func FromCoreList(data []review.ReviewCore) []ReviewResponse {
	var list []ReviewResponse
	for _, v := range data {
		list = append(list, FromCore(v))

	}
	return list
}
