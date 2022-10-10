package delivery

import (
	"capstone-project/config"
	"capstone-project/features/booking"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type bookingDelivery struct {
	bookingUsecase booking.UsecaseInterface
}

var event coreapi.Client

func New(e *echo.Echo, usecase booking.UsecaseInterface) {
	handler := &bookingDelivery{
		bookingUsecase: usecase,
	}

	e.POST("/bookings/addtocart", handler.PostData, middlewares.JWTMiddleware())
	e.GET("/bookings", handler.GetBooking, middlewares.JWTMiddleware())
	e.GET("/bookings/:id", handler.GetBookingId, middlewares.JWTMiddleware())
	e.POST("/bookings/:id/addpayment", handler.AddPayment, middlewares.JWTMiddleware())
	e.DELETE("bookings/:id", handler.DeleteBooking, middlewares.JWTMiddleware())
	e.POST("callback", handler.PaymentWebHook)
	e.GET("/history", handler.History, middlewares.JWTMiddleware())

}

func (delivery *bookingDelivery) PostData(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}

	var bookingDataRequest BookingRequest
	bookingDataRequest.UserID = uint(userId)
	// fmt.Println(bookingDataRequest.UserID)

	errBind := c.Bind(&bookingDataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}
	// fmt.Println(bookingDataRequest.)
	row, err := delivery.bookingUsecase.PostData(ToCore(bookingDataRequest))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input"))

}

func (delivery *bookingDelivery) GetBooking(c echo.Context) error {

	field_id, err := strconv.Atoi(c.QueryParam("field_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp(err.Error()))
	}
	user_id, errUser := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp(errUser.Error()))
	}

	venue_id, errVenue := strconv.Atoi(c.QueryParam("venue_id"))
	if err != nil && venue_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp(errVenue.Error()))
	}

	data, err := delivery.bookingUsecase.GetAllBooking(user_id, field_id, venue_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all data success", FromCoreList(data)))
}

func (delivery *bookingDelivery) History(c echo.Context) error {

	field_id, err := strconv.Atoi(c.QueryParam("field_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp(err.Error()))
	}
	user_id, errUser := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp(errUser.Error()))
	}

	data, err := delivery.bookingUsecase.History(user_id, field_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all data success", FromCoreList(data)))
}

func (delivery *bookingDelivery) GetBookingId(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.bookingUsecase.GetBookingById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}

func (delivery *bookingDelivery) AddPayment(c echo.Context) error {
	midtrans.ServerKey = config.MidtransServerKey()
	event.New(midtrans.ServerKey, midtrans.Sandbox)

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	dataBooking, errGetBooking := delivery.bookingUsecase.GetBookingById(id_conv)
	if errGetBooking != nil {
		return errGetBooking
	}

	var payment BookingRequest
	payment.FieldID = dataBooking.FieldID
	payment.UserID = dataBooking.UserID
	payment.ScheduleDetailID = dataBooking.ScheduleDetailID

	errBind := c.Bind(&payment)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}
	dataCore := ToCore(payment)
	dataCore.ID = uint(id_conv)
	dataCore.Total_price = dataBooking.Total_price

	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	timer := currentTime.Format("15:04:05")

	orderIDPay := fmt.Sprintf("Order-%s-%s-%s-%s", dataBooking.Name_User, dataBooking.Nama_venue, date, timer)
	dataCore.OrderID = orderIDPay

	inputPay := ToCoreMidtrans(dataCore)

	if payment.Payment_method == "BCA" {
		dataCore.Payment_method = "BCA"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	}

	if payment.Payment_method == "BRI" {
		dataCore.Payment_method = "BRI"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBri,
		}
	}

	if payment.Payment_method == "BNI" {
		dataCore.Payment_method = "BNI"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
		}
	}

	detailPay, errPay := delivery.bookingUsecase.CreatePaymentBankTransfer(int(dataCore.FieldID), int(dataCore.UserID), int(dataBooking.ScheduleDetailID), inputPay)

	if errPay != nil {
		return c.JSON(500, helper.Fail_Resp(errPay.Error()))
	}

	result := FromMidtransToPayment(detailPay)

	layout := "2006-01-02 15:04:05"
	trTime, _ := time.Parse(layout, detailPay.TransactionTime)
	result.TransactionTime = trTime
	result.TransactionExpire = trTime.Add(time.Hour * 24)

	dataCore.OrderID = result.OrderID
	dataCore.Status_payment = result.TransactionStatus
	dataCore.TransactionID = result.TransactionID
	dataCore.Virtual_account = result.BillNumber
	dataCore.Transaction_time = result.TransactionTime.String()
	dataCore.Transaction_exp = result.TransactionExpire.String()

	row, err := delivery.bookingUsecase.AddPayment(dataCore, id_conv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}

func (delivery *bookingDelivery) DeleteBooking(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)
	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	row, err := delivery.bookingUsecase.DeleteBooking(id_conv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail delete data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("rows affected 0, fail delete data"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success delete data"))
}

func (delivery *bookingDelivery) PaymentWebHook(c echo.Context) error {
	var webhookRequest MidtransHookRequest

	errBind := c.Bind(&webhookRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp(errBind.Error()))
	}

	err := delivery.bookingUsecase.PaymentWebHook(webhookRequest.OrderID, webhookRequest.TransactionStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success"))
}
