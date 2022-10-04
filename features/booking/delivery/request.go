package delivery

import "capstone-project/features/booking"

type BookingRequest struct {
	FieldID        uint   `json:"field_id" form:"field_id"`
	UserID         uint   `json:"user_id" form:"user_id"`
	Start_hours    uint   `json:"start_hours" form:"start_hours"`
	End_hours      uint   `json:"end_hours" form:"end_hours"`
	Payment_method string `json:"payment_method" form:"payment_method"`
}

func ToCore(data BookingRequest) booking.BookingCore {
	return booking.BookingCore{
		FieldID:        data.FieldID,
		UserID:         data.UserID,
		Start_hours:    data.Start_hours,
		End_hours:      data.End_hours,
		Payment_method: data.Payment_method,
	}

}
