package delivery

import (
	"capstone-project/features/booking"
)

type BookingResponse struct {
	ID               uint   `json:"booking_id"`
	Name_venue       string `json:"name_venue"`
	Name_User        string `json:"name_user"`
	field_id         uint   `json:"field_id"`
	Category         string `json:"category"`
	Start_hours      uint   `json:"start_hours"`
	End_hours        uint   `json:"end_hours"`
	Total_price      uint   `json:"total_price"`
	OrderID          uint   `json:"order_id"`
	Status_payment   string `json:"status_payment"`
	Transaction_time string `json:"transaction_time"`
}

func FromCore(data booking.BookingCore) BookingResponse {
	return BookingResponse{
		ID:               data.ID,
		Name_venue:       data.Nama_venue,
		Name_User:        data.Name_User,
		field_id:         data.FieldID,
		Category:         data.Category,
		Start_hours:      data.Start_hours,
		End_hours:        data.End_hours,
		Total_price:      data.Total_price,
		Status_payment:   data.Status_payment,
		Transaction_time: data.Transaction_time,
	}

}

func FromCoreList(data []booking.BookingCore) []BookingResponse {
	var list []BookingResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
