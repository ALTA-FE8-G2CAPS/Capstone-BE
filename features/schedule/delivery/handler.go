package delivery

import (
	"capstone-project/features/schedule"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type scheduleDelivery struct {
	scheduleUsecase schedule.UsecaseInterface
}

func New(e *echo.Echo, usecase schedule.UsecaseInterface) {
	handler := &scheduleDelivery{
		scheduleUsecase: usecase,
	}

	e.POST("/schedules", handler.PostSchedule, middlewares.JWTMiddleware())
	e.GET("/schedules", handler.GetSchedule, middlewares.JWTMiddleware())
	e.GET("/schedules/:id", handler.GetScheduleId, middlewares.JWTMiddleware())

}

func (delivery *scheduleDelivery) PostSchedule(c echo.Context) error {
	var scheduleRequestdata ScheduleRequest

	errBind := c.Bind(&scheduleRequestdata)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	row, err := delivery.scheduleUsecase.PostData(ToCore(scheduleRequestdata))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input"))
	}

	if row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input"))

}

func (delivery *scheduleDelivery) GetSchedule(c echo.Context) error {

	field_id, err := strconv.Atoi(c.QueryParam("field_id"))
	if err != nil && field_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail converse class_id param"))
	}

	data, err := delivery.scheduleUsecase.GetAllSchedule(field_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all data success", FromCoreList(data)))
}

func (delivery *scheduleDelivery) GetScheduleId(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.scheduleUsecase.GetScheduleById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}
