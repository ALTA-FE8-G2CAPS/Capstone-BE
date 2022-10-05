package delivery

import (
	"capstone-project/features/schedule"
)

type ScheduleResponse struct {
	ID          uint   `json:"schedule_id"`
	FieldID     uint   `json:"field_id"`
	Category    string `json:"category"`
	Day         string `json:"day"`
	Start_hours string `json:"start_hours"`
	End_hours   string `json:"end_hours"`
}

func FromCore(data schedule.ScheduleCore) ScheduleResponse {
	return ScheduleResponse{
		ID:          data.ID,
		FieldID:     data.FieldID,
		Category:    data.Category,
		Day:         data.Day,
		Start_hours: data.Start_hours,
		End_hours:   data.End_hours,
	}

}

func FromCoreList(data []schedule.ScheduleCore) []ScheduleResponse {
	var list []ScheduleResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
