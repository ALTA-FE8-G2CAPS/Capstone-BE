package venue

type VenueCore struct {
	ID           uint
	Name_Venue   string
	Addres_Venue string
	Description  string
	UserID       uint
	latitude     float64
	longitude    float64
	Foto_Venue   []FotoVenue
}

type FotoVenue struct {
	VenueID    uint
	Foto_Venue string
}

type UsecaseInterface interface {
	GetAllVenue() (data []VenueCore, err error)
	GetVenueById(id int) (data VenueCore, err error)
	PostData(data VenueCore) (row int, err error)
	PutData(data VenueCore) (row int, err error)
	DeleteVenue(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllVenue() (data []VenueCore, err error)
	SelectVenueById(id int) (data VenueCore, err error)
	InsertData(data VenueCore) (row int, err error)
	UpdateVenue(data VenueCore) (row int, err error)
	DeleteVenue(id int) (row int, err error)
}
