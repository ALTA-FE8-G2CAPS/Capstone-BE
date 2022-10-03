package venue

type VenueCore struct {
	ID                uint
	Name_venue        string
	Address_venue     string
	Description_venue string
	UserID            uint
	Name_user         string
	Latitude          float64
	Longitude         float64
}

type FotoVenue struct {
	VenueID    uint
	Foto_Venue string
}

type UsecaseInterface interface {
	GetAllVenue(user_id int) (data []VenueCore, err error)
	GetVenueById(id int) (data VenueCore, err error)
	PostData(data VenueCore) (row int, err error)
	PutData(data VenueCore, user_id int) (row int, err error)
	DeleteVenue(user_id, venue_id int) (row int, err error)
}

type DataInterface interface {
	SelectAllVenue(user_id int) (data []VenueCore, err error)
	SelectVenueById(id int) (data VenueCore, err error)
	InsertData(data VenueCore) (row int, err error)
	UpdateVenue(data VenueCore, user_id int) (row int, err error)
	DeleteVenue(user_id, venue_id int) (row int, err error)
}
