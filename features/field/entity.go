package field

type FieldCore struct {
	ID         uint
	VenueID    uint
	Name_venue string
	Category   string
	Price      uint
	Schedule   Schedule
}

type Schedule struct {
	ID              uint
	FieldID         uint
	Day             string
	Start_hours     uint
	End_hours       uint
	SchedulesDetail []ScheduleDetail
}

type ScheduleDetail struct {
	ID              uint
	ScheduleID      uint
	Start_hours     uint
	End_hours       uint
	Status_schedule string
}

type UsecaseInterface interface {
	GetAllField(venue_id int) (data []FieldCore, er error)
	GetFieldById(id int) (data FieldCore, err error)
	PostData(data FieldCore) (row int, err error)
	PutData(data FieldCore, user_id int) (row int, err error)
	DeleteField(user_id, field_id int) (row int, err error)
}

type DataInterface interface {
	SelectAllField(venue_id int) (data []FieldCore, err error)
	SelectFieldById(id int) (data FieldCore, err error)
	InsertData(data FieldCore) (row int, err error)
	UpdateField(data FieldCore, user_id int) (row int, err error)
	DeleteField(user_id, field_id int) (row int, err error)
}
