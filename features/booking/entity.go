package booking

type BookingCore struct {
	ID               uint
	UserID           uint
	Name_User        string
	Nama_venue       string
	FieldID          uint
	Category         string
	Start_hours      uint
	End_hours        uint
	Total_price      uint
	Payment_method   string
	OrderID          uint
	TransactionID    uint
	Status_payment   string
	Virtual_account  string
	Transaction_time string
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
	GetAllBooking(user_id, field_id int) (data []BookingCore, err error)
	GetBookingById(id int) (data BookingCore, err error)
	PostData(data BookingCore) (row int, err error)
	// PutData(data BookingCore, user_id int) (row int, err error)
	// DeleteBooking(user_id, Booking_id int) (row int, err error)
	// PostPhoto(data FotoBooking) (row int, err error)
}

type DataInterface interface {
	SelectAllBooking(user_id, field_id int) (data []BookingCore, err error)
	SelectBookingById(id int) (data BookingCore, err error)
	InsertData(data BookingCore) (row int, err error)
	// UpdateBooking(data BookingCore, user_id int) (row int, err error)
	// DeleteBooking(user_id, Booking_id int) (row int, err error)
	// UploadPhoto(data FotoBooking) (row int, err error)
}
