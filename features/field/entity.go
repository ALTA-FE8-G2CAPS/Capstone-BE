package field

type FieldCore struct {
	ID       uint
	VenueID  uint
	Category string
	Price    int
}

type UsecaseInterface interface {
	GetAllField(venue_id int) (data []FieldCore, er error)
	// GetFieldById(id int) (data FieldCore, err error)
	// PostData(data FieldCore) (row int, err error)
	// PutData(data FieldCore) (row int, err error)
	// DeleteField(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllField() (data []FieldCore, err error)
	// SelectFieldById(id int) (data FieldCore, err error)
	// InserData(data FieldCore) (row int, err error)
	// UpdateField(data FieldCore) (row int, err error)
	// DeleteField(id int) (row int, err error)
}
