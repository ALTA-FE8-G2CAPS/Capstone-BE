package delivery

import (
	"capstone-project/config"
	"capstone-project/features/payment"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type paymentDelivery struct {
	paymentUsecase payment.UsecaseInterface
}

func New(usecase payment.UsecaseInterface) *paymentDelivery {
	return &paymentDelivery{
		paymentUsecase: usecase,
	}
}

var booking coreapi.Client

func (delivery *paymentDelivery) PostPayment(c echo.Context) error {
	midtrans.ServerKey = config.MidtransOrderServerKey()
	booking.New(midtrans.ServerKey, midtrans.Sandbox)
	typeName := c.Param("type")
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail extract token"))
	}

	bookingOrder, totalPrice, err := delivery.paymentUsecase.CreatePayment(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail operation"))
	}
	var Transfer PaymentRequest
	switch {
	case typeName == "bni":
		Transfer = PaymentRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(bookingOrder),
				GrossAmt: int64(totalPrice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
		}
	case typeName == "bca":
		Transfer = PaymentRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(bookingOrder),
				GrossAmt: int64(totalPrice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
		}
	case typeName == "bri":
		Transfer = PaymentRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(bookingOrder) + "bri",
				GrossAmt: int64(totalPrice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
		}
	case typeName == "permata":
		Transfer = PaymentRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(bookingOrder) + "permata1",
				GrossAmt: int64(totalPrice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankPermata,
				Permata: &coreapi.PermataBankTransferDetail{
					RecipientName: "lamiapp",
				},
			},
		}
	case typeName == "mandiri":
		Transfer = PaymentRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(bookingOrder),
				GrossAmt: int64(totalPrice),
			},
		}
	default:
		return c.JSON(http.StatusInternalServerError, "fail pokoknya fail")
	}
	var TransferCore coreapi.ChargeReq
	if typeName == "bni" || typeName == "bca" || typeName == "bri" {
		TransferCore = ToCore(Transfer)
	}
	resp, err := coreapi.ChargeTransaction(&TransferCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail operationing"))
	}
	if typeName == "bni" || typeName == "bca" || typeName == "bri" || typeName == "permata" {
		return c.JSON(http.StatusOK, FromCoreChargeMidtrans(*resp))
	}
	return c.JSON(http.StatusOK, FromCoreChargeMidtrans(*resp))
}
