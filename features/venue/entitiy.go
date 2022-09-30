package venue

type VenueCore struct {
	ID                uint
	Name_Venue        string
	Address_Venue     string
	Description_venue string
	Nama_user         string
	Latitude          float64
	Longitude         float64
	Foto_Venue        []FotoVenue
}

type FotoVenue struct {
	VenueID    uint
	Foto_Venue string
}

type UsecaseInterface interface {
	GetAllVenue(user_id int) (data []VenueCore, err error)
	GetVenueById(id int) (data VenueCore, err error)
	PostData(data VenueCore) (row int, err error)
	PutData(data VenueCore) (row int, err error)
	DeleteVenue(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllVenue(user_id int) (data []VenueCore, err error)
	SelectVenueById(id int) (data VenueCore, err error)
	InsertData(data VenueCore) (row int, err error)
	UpdateVenue(data VenueCore) (row int, err error)
	DeleteVenue(id int) (row int, err error)
}
