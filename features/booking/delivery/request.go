package delivery

import (
	"capstone-project/features/booking"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type BookingRequest struct {
	FieldID          uint   `json:"field_id" form:"field_id"`
	UserID           uint   `json:"user_id" form:"user_id"`
	ScheduleDetailID uint   `json:"schedule_detail_id" form:"schedule_detail_id"`
	Payment_method   string `json:"payment_method" form:"payment_method"`
	OrderID          string `json:"order_id" form:"order_id"`
	TransactionID    string `json:"transaction_id" form:"transaction_id"`
	Virtual_account  string `json:"virtual_account" form:"virtual_account"`
	Transaction_time string `json:"transaction_time" form:"transaction_time"`
	Transaction_exp  string `json:"transaction_exp" form:"transaction_exp"`
}

func ToCore(data BookingRequest) booking.BookingCore {
	return booking.BookingCore{
		FieldID:          data.FieldID,
		UserID:           data.UserID,
		ScheduleDetailID: data.ScheduleDetailID,
		Payment_method:   data.Payment_method,
		OrderID:          data.OrderID,
		TransactionID:    data.TransactionID,
		Virtual_account:  data.Virtual_account,
		Transaction_time: data.Transaction_time,
		Transaction_exp:  data.Transaction_exp,
	}

}

func ToCoreMidtrans(req booking.BookingCore) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: int64(req.Total_price),
		},
	}
}

type MidtransHookRequest struct {
	TransactionTime   string `form:"transaction_time" json:"transaction_time"`
	TransactionStatus string `form:"transaction_status" json:"transaction_status"`
	OrderID           string `form:"order_id" json:"order_id"`
	MerchantID        string `form:"merchant_id" json:"merchant_id"`
	GrossAmount       string `form:"gross_amount" json:"gross_amount"`
	FraudStatus       string `form:"fraud_status" json:"fraud_status"`
	Currency          string `form:"currency" json:"currency"`
}
