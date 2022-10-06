package delivery

import "github.com/midtrans/midtrans-go/coreapi"

type PaymentResponse struct {
	TransactionTime   string
	TransactionStatus string
	PaymentType       string
	VAnumbers         VAnumbers
	OrderID           string
	GroosAmt          string
}

type VAnumbers struct {
	BankTransfer string
	VAnumber     string
}

type TransactionStatusResponse struct {
	OrderID           string
	TransactionTime   string
	TransactionStatus string
	Bank              string
	GrossAmount       string
	PaymentType       string
	SettlementTime    string
}

func FromCoreChargeMidtrans(resp coreapi.ChargeResponse) PaymentResponse {
	return PaymentResponse{
		TransactionTime:   resp.TransactionTime,
		TransactionStatus: resp.TransactionStatus,
		PaymentType:       resp.PaymentType,
		VAnumbers: VAnumbers{
			BankTransfer: resp.VaNumbers[0].Bank,
			VAnumber:     resp.VaNumbers[0].VANumber,
		},
		OrderID:  resp.OrderID,
		GroosAmt: resp.GrossAmount,
	}
}

func FromCoreStatusResponse(data coreapi.TransactionStatusResponse) TransactionStatusResponse {
	return TransactionStatusResponse{
		OrderID:           data.OrderID,
		TransactionTime:   data.TransactionTime,
		TransactionStatus: data.TransactionStatus,
		GrossAmount:       data.GrossAmount,
		PaymentType:       data.PaymentType,
		SettlementTime:    data.SettlementTime,
	}
}
