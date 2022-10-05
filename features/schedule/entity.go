package schedule

type ScheduleCore struct {
	ID          uint
	FieldID     uint
	Category    string
	Day         string
	Start_hours uint
	End_hours   uint
}

type UsecaseInterface interface {
	GetAllField(field_id int) (data []ScheduleCore, err error)
	GetFieldById(id int) (data ScheduleCore, err error)
	PostData(data ScheduleCore) (row int, err error)
	// PutData(data FieldCore, user_id int) (row int, err error)
	// DeleteField(user_id, field_id int) (row int, err error)
}

type DataInterface interface {
	SelectAllField(field_id int) (data []ScheduleCore, err error)
	SelectFieldById(id int) (data ScheduleCore, err error)
	InsertData(data ScheduleCore) (row int, err error)
	// UpdateField(data FieldCore, user_id int) (row int, err error)
	// DeleteField(user_id, field_id int) (row int, err error)
}
