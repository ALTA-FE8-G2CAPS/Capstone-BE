package delivery

import "capstone-project/features/schedule"

type ScheduleRequest struct {
	FieldID     uint   `json:"field_id" form:"field_id"`
	Day         string `json:"day" form:"day"`
	Start_hours string `json:"start_hours" form:"start_hours"`
	End_hours   string `json:"end_hours" form:"end_hours"`
}

type ScheduleDetailRequest struct {
	ScheduleID      uint   `json:"schedule_id" form:"schedule_id"`
	Start_hours     string `json:"start_hours" form:"start_hours"`
	End_hours       string `json:"end_hours" form:"end_hours"`
	Status_schedule string `json:"status_schedule" form:"status_schedule"`
}

func ToCore(data ScheduleRequest) schedule.ScheduleCore {
	return schedule.ScheduleCore{
		FieldID:     data.FieldID,
		Day:         data.Day,
		Start_hours: data.Start_hours,
		End_hours:   data.End_hours,
	}
}

func ToCoreScheduleDetail(data ScheduleDetailRequest) schedule.ScheduleDetailCore {
	return schedule.ScheduleDetailCore{
		ScheduleID:      data.ScheduleID,
		Start_hours:     data.Start_hours,
		End_hours:       data.End_hours,
		Status_schedule: data.Status_schedule,
	}
}
