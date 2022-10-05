package delivery

import "capstone-project/features/schedule"

type ScheduleRequest struct {
	FieldID     uint   `json:"field_id" form:"field_id"`
	Day         string `json:"day" form:"day"`
	Start_hours string `json:"start_hours" form:"start_hours"`
	End_hours   string `json:"end_hours" form:"end_hours"`
}

func ToCore(data ScheduleRequest) schedule.ScheduleCore {
	return schedule.ScheduleCore{
		FieldID:     data.FieldID,
		Day:         data.Day,
		Start_hours: data.Start_hours,
		End_hours:   data.End_hours,
	}
}
