package field

type FieldCore struct {
	ID         uint
	VenueID    uint
	Name_venue string
	Category   string
	Price      float64
}

type UsecaseInterface interface {
	GetAllField(venue_id int) (data []FieldCore, er error)
	GetFieldById(id int) (data FieldCore, err error)
	PostData(data FieldCore) (row int, err error)
	PutData(data FieldCore) (row int, err error)
	DeleteField(venue_id, field_id int) (row int, err error)
}

type DataInterface interface {
	SelectAllField(venue_id int) (data []FieldCore, err error)
	SelectFieldById(id int) (data FieldCore, err error)
	InsertData(data FieldCore) (row int, err error)
	UpdateField(data FieldCore) (row int, err error)
	DeleteField(vneue_id, field_id int) (row int, err error)
}
