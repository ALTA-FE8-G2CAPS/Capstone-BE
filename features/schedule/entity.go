package schedule

type ScheduleCore struct {
	ID             uint
	FieldID        uint
	Category       string
	Day            string
	Start_hours    string
	End_hours      string
	ScheduleDetail []ScheduleDetailCore
}

type ScheduleDetailCore struct {
	ID              uint
	ScheduleID      uint
	Start_hours     string
	End_hours       string
	Status_schedule string
}

type UsecaseInterface interface {
	GetAllSchedule(field_id int) (data []ScheduleCore, err error)
	GetScheduleById(id int) (data ScheduleCore, err error)
	PostData(dataSchedule ScheduleCore) (row int, err error)
	// PutData(data FieldCore, user_id int) (row int, err error)
	// DeleteField(user_id, field_id int) (row int, err error)
}

type DataInterface interface {
	SelectAllSchedule(field_id int) (data []ScheduleCore, err error)
	SelectScheduleById(id int) (data ScheduleCore, err error)
	InsertData(dataSchedule ScheduleCore) (shcedule_id, row int, err error)
	InsertDetailSchedule(schedule_id int, dataGenerete []map[string]interface{}) (row int, err error)
	// UpdateField(data FieldCore, user_id int) (row int, err error)
	// DeleteField(user_id, field_id int) (row int, err error)
}
