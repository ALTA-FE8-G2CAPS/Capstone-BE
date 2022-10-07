package usecase

import (
	"capstone-project/features/booking"
	field "capstone-project/features/field"

	"errors"
)

type BookingUsecase struct {
	bookingData booking.DataInterface
	fieldData   field.DataInterface
}

func New(dataBooking booking.DataInterface, dataField field.DataInterface) booking.UsecaseInterface {
	return &BookingUsecase{
		bookingData: dataBooking, // fieldData
		fieldData:   dataField,
	}
}

func (usecase *BookingUsecase) GetAllBooking(user_id, field_id int) ([]booking.BookingCore, error) {
	dataField, err := usecase.bookingData.SelectAllBooking(user_id, field_id)
	if err != nil {
		return nil, err
	}
	return dataField, nil
}

func (usecase *BookingUsecase) PostData(data booking.BookingCore) (row int, err error) {
	var newBooking booking.BookingCore
	if data.FieldID == 0 || data.Start_hours == 0 || data.End_hours == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	dataField, errField := usecase.fieldData.SelectFieldById(int(data.FieldID))
	if errField != nil {
		return -1, errField
	}
	newBooking.UserID = data.UserID
	newBooking.Name_User = data.Name_User
	newBooking.Nama_venue = dataField.Name_venue
	newBooking.FieldID = data.FieldID
	newBooking.Category = data.Category
	newBooking.Start_hours = data.Start_hours
	newBooking.End_hours = data.End_hours
	newBooking.Total_price = ((data.End_hours - data.Start_hours) * dataField.Price)
	newBooking.Payment_method = data.Payment_method
	newBooking.TransactionID = data.TransactionID
	newBooking.Status_payment = data.Status_payment
	newBooking.Virtual_account = data.Virtual_account
	newBooking.Transaction_time = data.Transaction_time

	row, err = usecase.bookingData.InsertData(newBooking)
	if err != nil {
		return -1, err
	}

	// helper.SendGmailNotif("muhammadadityogunawan@gmail.com", "lapangan volly", "gor bung karno", "50000", "1", "50000", "50000", "50000", "0")
	return row, err
}

func (usecase *BookingUsecase) GetBookingById(id int) (booking.BookingCore, error) {
	result, err := usecase.bookingData.SelectBookingById(id)
	if err != nil {
		return booking.BookingCore{}, errors.New("data tidak ditemukan")
	}
	return result, nil
}
