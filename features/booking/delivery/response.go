package delivery

import (
	"capstone-project/config"
	"capstone-project/features/booking"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type BookingResponse struct {
	ID               uint   `json:"booking_id"`
	VenueID          uint   `json:"venue_id"`
	Name_venue       string `json:"name_venue"`
	Name_User        string `json:"name_user"`
	Email            string `json:"email"`
	Field_id         uint   `json:"field_id"`
	Category         string `json:"category"`
	ScheduleDetailID uint   `json:"schedule_detail_id"`
	Start_hours      string `json:"start_hours"`
	End_Hours        string `json:"end_hours"`
	Price            uint   `json:"price"`
	Total_price      uint   `json:"total_price"`
	Payment_method   string `json:"payment_method"`
	OrderID          string `json:"order_id"`
	Status_payment   string `json:"status_payment"`
	Transaction_time string `json:"transaction_time"`
	TransactionID    string `json:"transaction_id" form:"transaction_id"`
	Transaction_exp  string `json:"transaction_exp"`
	Virtual_account  string `json:"virtual_account" form:"virtual_account"`
}

type Payment struct {
	OrderID           string    `json:"orderID" form:"orderID"`
	TransactionID     string    `json:"transactionID" form:"transactionID"`
	PaymentMethod     string    `json:"paymentMethod" form:"paymentMethod"`
	BillNumber        string    `json:"billNumber" form:"billNumber"`
	Bank              string    `json:"bank" form:"bank"`
	GrossAmount       string    `json:"grossAmount" form:"grossAmount"`
	TransactionTime   time.Time `json:"transactionTime" form:"transactionTime"`
	TransactionExpire time.Time `json:"transactionExpired" form:"transactionExpired"`
	TransactionStatus string    `json:"transactionStatus" form:"transactionStatus"`
}

func FromMidtransToPayment(resMidtrans *coreapi.ChargeResponse) Payment {
	return Payment{
		OrderID:           resMidtrans.OrderID,
		TransactionID:     resMidtrans.TransactionID,
		PaymentMethod:     config.PaymentBankTransferBCA,
		BillNumber:        resMidtrans.VaNumbers[0].VANumber,
		Bank:              resMidtrans.VaNumbers[0].Bank,
		GrossAmount:       resMidtrans.GrossAmount,
		TransactionStatus: resMidtrans.TransactionStatus,
	}
}

func FromCore(data booking.BookingCore) BookingResponse {
	return BookingResponse{
		ID:               data.ID,
		VenueID:          data.VenueID,
		Name_venue:       data.Nama_venue,
		Name_User:        data.Name_User,
		Email:            data.Email,
		Field_id:         data.FieldID,
		Category:         data.Category,
		ScheduleDetailID: data.ScheduleDetailID,
		Start_hours:      data.Start_hours,
		End_Hours:        data.End_Hours,
		Price:            data.Price,
		Total_price:      data.Total_price,
		Payment_method:   data.Payment_method,
		OrderID:          data.OrderID,
		Status_payment:   data.Status_payment,
		Transaction_time: data.Transaction_time,
		Transaction_exp:  data.Transaction_exp,
		TransactionID:    data.TransactionID,
		Virtual_account:  data.Virtual_account,
	}

}

func FromCoreList(data []booking.BookingCore) []BookingResponse {
	var list []BookingResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
