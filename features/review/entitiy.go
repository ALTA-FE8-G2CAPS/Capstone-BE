package review

type ReviewCore struct {
	ID          uint
	Name_user   string
	Foto_user   string
	UserID      uint
	Nama_venue  string
	VenueID     uint
	Rate        uint
	Feedback    string
	Foto_review string
}

type Usecaseinterface interface {
	PostReview(data ReviewCore) (row int, err error)
	GetReviewById(venue_id int) (data []ReviewCore, err error)
}

type DataInterface interface {
	InsertReview(data ReviewCore) (row int, err error)
	SelectReviewById(venue_id int) (data []ReviewCore, err error)
}
