package delivery

import (
	"capstone-project/features/booking"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookingDelivery struct {
	bookingUsecase booking.UsecaseInterface
}

func New(e *echo.Echo, usecase booking.UsecaseInterface) {
	handler := &bookingDelivery{
		bookingUsecase: usecase,
	}

	e.POST("/bookings/addtocart", handler.PostData, middlewares.JWTMiddleware())
	e.GET("/bookings", handler.GetBooking, middlewares.JWTMiddleware())
	e.GET("/bookings/:id", handler.GetBookingId, middlewares.JWTMiddleware())

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

	row, err := delivery.bookingUsecase.PostData(ToCore(bookingDataRequest))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input"))

}

func (delivery *bookingDelivery) GetBooking(c echo.Context) error {

	field_id, err := strconv.Atoi(c.QueryParam("field_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail converse class_id param"))
	}
	user_id, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail converse class_id param"))
	}

	data, err := delivery.bookingUsecase.GetAllBooking(user_id, field_id)

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
