package helper

import (
	"fmt"
	"time"
)

func GenerateSchedule(shcedule_id int, start_hours, end_hours string) []map[string]interface{} {
	var result = []map[string]interface{}{}
	layoutFormat := "15:04"
	dateStart, _ := time.Parse(layoutFormat, start_hours)
	dateEnd, _ := time.Parse(layoutFormat, end_hours)
	for t := dateStart; t.Before(dateEnd); t = t.Add(1 * time.Hour) {
		var hasil = map[string]interface{}{}
		dateFormatStart := t.Format(layoutFormat)
		hasil["status_schedule"] = "Available"
		hasil["schedule_id"] = fmt.Sprintf("%d", shcedule_id)
		hasil["start_hours"] = dateFormatStart
		timeEnd := t.Add(1 * time.Hour)
		hasil["end_hours"] = timeEnd.Format(layoutFormat)
		result = append(result, hasil)

	}
	return result
}
