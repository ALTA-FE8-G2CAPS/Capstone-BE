package usecase

import (
	"capstone-project/features/booking"
	field "capstone-project/features/field"
	schedule "capstone-project/features/schedule"
	"capstone-project/utils/helper"

	"errors"

	"github.com/midtrans/midtrans-go/coreapi"
)

type BookingUsecase struct {
	bookingData  booking.DataInterface
	fieldData    field.DataInterface
	scheduleData schedule.DataInterface
}

func New(dataBooking booking.DataInterface, dataField field.DataInterface, dataSchedule schedule.DataInterface) booking.UsecaseInterface {
	return &BookingUsecase{
		bookingData:  dataBooking, // fieldData
		fieldData:    dataField,
		scheduleData: dataSchedule,
	}
}

func (usecase *BookingUsecase) GetAllBooking(user_id, field_id, venue_id int) ([]booking.BookingCore, error) {
	dataField, err := usecase.bookingData.SelectAllBooking(user_id, field_id, venue_id)
	if err != nil {
		return nil, err
	}
	return dataField, nil
}
func (usecase *BookingUsecase) History(user_id, field_id int) ([]booking.BookingCore, error) {
	dataField, err := usecase.bookingData.History(user_id, field_id)
	if err != nil {
		return nil, err
	}
	return dataField, nil
}

func (usecase *BookingUsecase) PostData(data booking.BookingCore) (row int, err error) {
	var newBooking booking.BookingCore
	if data.FieldID == 0 || data.ScheduleDetailID == 0 || data.VenueID == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	// dataSchedule, _ := usecase.scheduleData.SelectScheduleDetailById(int(newBooking.ScheduleDetailID))
	// if dataSchedule.Status_schedule != "Available" {
	// 	return -1, errors.New("field already booked")
	// }

	dataField, errField := usecase.fieldData.SelectFieldById(int(data.FieldID))
	if errField != nil {
		return -1, errField
	}
	newBooking.UserID = data.UserID
	newBooking.Name_User = data.Name_User
	newBooking.Email = data.Email
	newBooking.VenueID = data.VenueID
	newBooking.Nama_venue = dataField.Name_venue
	newBooking.FieldID = data.FieldID
	newBooking.Category = data.Category
	newBooking.ScheduleDetailID = data.ScheduleDetailID
	newBooking.Start_hours = data.Start_hours
	newBooking.End_Hours = data.End_Hours
	newBooking.Price = data.Price
	newBooking.Total_price = dataField.Price
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

func (usecase *BookingUsecase) AddPayment(data booking.BookingCore, booking_id int) (int, error) {
	row, err := usecase.bookingData.UpdatePayment(data, booking_id)
	if err != nil {
		return -1, err
	}
	dataBooking, _ := usecase.bookingData.SelectBookingById(booking_id)
	helper.SendGmailNotif(dataBooking.Email, dataBooking.Name_User, dataBooking.Category, dataBooking.OrderID, dataBooking.Nama_venue, dataBooking.Total_price, 1, int(dataBooking.Total_price), dataBooking.Price)
	return row, nil
}

func (usecase *BookingUsecase) DeleteBooking(booking_id int) (row int, err error) {
	result, err := usecase.bookingData.DeleteBooking(booking_id)
	if err != nil {
		return -1, err
	}
	return result, err
}

func (usecase *BookingUsecase) CreatePaymentBankTransfer(field_id, user_id, schedule_detail_id int, reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	createPay, errCreatePay := usecase.bookingData.CreateDataPayment(reqPay)
	if errCreatePay != nil {
		return nil, errors.New("failed get response payment")
	}

	dataScheduleDetail, _ := usecase.scheduleData.SelectScheduleDetailById(schedule_detail_id)
	dataScheduleDetail.Status_schedule = "booked"

	usecase.scheduleData.UpdateScheduleDetail(dataScheduleDetail, schedule_detail_id)
	return createPay, nil
}

func (usecase *BookingUsecase) PaymentWebHook(OrderID, status string) error {
	var PaymentCore booking.BookingCore

	PaymentCore.OrderID = OrderID
	switch status {
	case "settlement":
		PaymentCore.Status_payment = "Paid"
	case "pending":
		PaymentCore.Status_payment = "Pending"
	default:
		PaymentCore.Status_payment = "Canceled"
	}

	_, err := usecase.bookingData.UpdatepaymentWebhook(PaymentCore)
	if err != nil {
		return err
	}

	return nil
}
