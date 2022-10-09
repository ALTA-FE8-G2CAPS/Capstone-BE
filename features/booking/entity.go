package booking

import "github.com/midtrans/midtrans-go/coreapi"

// import "github.com/midtrans/midtrans-go/coreapi"

type BookingCore struct {
	ID               uint
	UserID           uint
	Name_User        string
	VenueID          uint
	Nama_venue       string
	Email            string
	FieldID          uint
	Category         string
	ScheduleDetailID uint
	Start_hours      string
	End_Hours        string
	Total_price      uint
	Payment_method   string
	OrderID          string
	TransactionID    string
	Status_payment   string
	Virtual_account  string
	Transaction_time string
	Transaction_exp  string
	Field            Field
}

type Field struct {
	FieldID    uint
	VenueID    uint
	Name_venue string
	Category   string
	Price      float64
}

type UsecaseInterface interface {
	CreatePaymentBankTransfer(field_id, user_id, schedule_detail_id int, reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	GetAllBooking(user_id, field_id, venue_id int) (data []BookingCore, err error)
	GetBookingById(id int) (data BookingCore, err error)
	PostData(data BookingCore) (row int, err error)
	AddPayment(data BookingCore, booking_id int) (row int, err error)
	DeleteBooking(booking_id int) (row int, err error)
	PaymentWebHook(orderID, TransactionStatus string) error
	// PostPhoto(data FotoBooking) (row int, err error)
}

type DataInterface interface {
	CreateDataPayment(reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error)

	SelectAllBooking(user_id, field_id, venue_id int) (data []BookingCore, err error)
	SelectBookingById(id int) (data BookingCore, err error)
	InsertData(data BookingCore) (row int, err error)
	UpdatePayment(data BookingCore, booking_id int) (row int, err error)
	UpdatepaymentWebhook(data BookingCore) (int, error)
	DeleteBooking(booking_id int) (row int, err error)
	// SelectPayment(orderID string) (BookingCore, error)

	// UploadPhoto(data FotoBooking) (row int, err error)
}
