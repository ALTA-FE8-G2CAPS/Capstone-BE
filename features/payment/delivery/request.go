package delivery

import (
	"capstone-project/features/user/data"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentRequest struct {
	PaymentType        coreapi.CoreapiPaymentType
	TransactionDetails midtrans.TransactionDetails
	BankTransfer       *coreapi.BankTransferDetails
	Booking            data.Booking
}

func ToCore(req PaymentRequest) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.TransactionDetails.OrderID,
			GrossAmt: req.TransactionDetails.GrossAmt,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: req.BankTransfer.Bank,
		},
	}
}
