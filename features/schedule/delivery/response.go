package delivery

import (
	"capstone-project/features/schedule"
)

type ScheduleResponse struct {
	ID             uint             `json:"schedule_id"`
	VenueID        uint             `json:"venue_id"`
	FieldID        uint             `json:"field_id"`
	Category       string           `json:"category"`
	Day            string           `json:"day"`
	Start_hours    string           `json:"start_hours"`
	End_hours      string           `json:"end_hours"`
	ScheduleDetail []ScheduleDetail `json:"schedule_detail"`
}
type ScheduleDetail struct {
	ID              uint   `json:"schedule_detail_id"`
	ScheduleID      uint   `json:"schedule_id"`
	VenueID         uint   `json:"venue_id"`
	Start_hours     string `json:"start_hours"`
	End_hours       string `json:"end_hours"`
	Status_schedule string `json:"status_schedule"`
}

func FromCore(data schedule.ScheduleCore) ScheduleResponse {
	return ScheduleResponse{
		ID:             data.ID,
		VenueID:        data.VenueID,
		FieldID:        data.FieldID,
		Category:       data.Category,
		Day:            data.Day,
		Start_hours:    data.Start_hours,
		End_hours:      data.End_hours,
		ScheduleDetail: FromCoreScheduleDetailList(data.ScheduleDetail),
	}

}
func FromCoreScheduleDetail(data schedule.ScheduleDetailCore) ScheduleDetail {
	return ScheduleDetail{
		ID:              data.ID,
		VenueID:         data.VenueID,
		ScheduleID:      data.ScheduleID,
		Start_hours:     data.Start_hours,
		End_hours:       data.End_hours,
		Status_schedule: data.Status_schedule,
	}

}

func FromCoreList(data []schedule.ScheduleCore) []ScheduleResponse {
	var list []ScheduleResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}

func FromCoreScheduleDetailList(data []schedule.ScheduleDetailCore) []ScheduleDetail {
	var list []ScheduleDetail
	for _, v := range data {
		var scheduleDetail ScheduleDetail
		scheduleDetail.ID = v.ID
		scheduleDetail.VenueID = v.VenueID
		scheduleDetail.ScheduleID = v.ScheduleID
		scheduleDetail.Start_hours = v.Start_hours
		scheduleDetail.End_hours = v.End_hours
		scheduleDetail.Status_schedule = v.Status_schedule
		list = append(list, scheduleDetail)
	}
	return list
}
