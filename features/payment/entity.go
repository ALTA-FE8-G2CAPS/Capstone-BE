package payment

import (
	"capstone-project/features/booking/data"
)

type PaymentCore struct {
	PaymentType        string
	Fields             FieldDetails
	User               UserDetails
	TransactionDetails TransactionDetails
	BankTransfer       BankTransferDetails
	Gopay              GopayDetails
	Shopeepay          ShopeepayDetails
	Booking            data.Booking
}

type FieldDetails struct {
	ID          uint
	Category    string
	Price       uint
	Start_hours string
	End_hours   string
	Nama_venue  string
}

type UserDetails struct {
	Name_user string
	Email     string
}

type TransactionDetails struct {
	OrderID  string
	GrossAmt int
}

type BankTransferDetails struct {
	BankName       string
	VANumber       string
	BCA            BCABankTransferDetail
	LangIDInquiry  string
	LangENInquiry  string
	LangIDPayment  string
	LangENPayment  string
	SubCompanyCode string
	BillInfo1      string
	BillInfo2      string
	BillInfoKey    string
}

type BCABankTransferDetail struct {
	RecipientName string
}

type GopayDetails struct {
	EnableCallback     bool
	CallbackUrl        string
	AccountID          string
	PaymentOptionToken string
}

type ShopeepayDetails struct {
	CallbackUrl string
}

type VAnumbers struct {
	BankTransfer string
	VAnumber     string
}

type DataInterface interface {
	DataPayment(userId int) (int, int, error)
}

type UsecaseInterface interface {
	CreatePayment(idUser int) (int, int, error)
}
