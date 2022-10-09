package schedule

type ScheduleCore struct {
	ID             uint
	FieldID        uint
	VenueID        uint
	Category       string
	Day            string
	Start_hours    string
	End_hours      string
	ScheduleDetail []ScheduleDetailCore
}

type ScheduleDetailCore struct {
	ID              uint
	ScheduleID      uint
	VenueID         uint
	Start_hours     string
	End_hours       string
	Status_schedule string
}

type UsecaseInterface interface {
	GetAllSchedule(field_id int) (data []ScheduleCore, err error)
	GetScheduleById(id int) (data ScheduleCore, err error)
	GetScheduleDetailById(id int) (data ScheduleDetailCore, err error)
	PostData(dataSchedule ScheduleCore) (row int, err error)
	PutScheduleDetail(data ScheduleDetailCore, id int) (row int, err error)
	PutData(data ScheduleCore, schedule_id int) (row int, err error)
	DeleteSchedule(schedule_id int) (row int, err error)
}

type DataInterface interface {
	SelectAllSchedule(field_id int) (data []ScheduleCore, err error)
	SelectScheduleById(id int) (data ScheduleCore, err error)
	SelectScheduleDetailById(id int) (data ScheduleDetailCore, err error)
	InsertData(dataSchedule ScheduleCore) (shcedule_id, row int, err error)
	InsertDetailSchedule(schedule_id int, dataGenerete []map[string]interface{}) (row int, err error)
	UpdateScheduleDetail(data ScheduleDetailCore, id int) (row int, err error)
	UpdateSchedule(data ScheduleCore, schedule_id int) (row int, err error)
	DeleteSchedule(schedule_id int) (row int, err error)
}
