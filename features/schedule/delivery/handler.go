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
	e.GET("/schedulesdetail/:id", handler.GetScheduleDetail, middlewares.JWTMiddleware())
	e.PUT("schedulesdetail/:id", handler.UpdateScheduleDetail, middlewares.JWTMiddleware())
	e.DELETE("/schedules/:id", handler.DeleteSchedule, middlewares.JWTMiddleware())
	e.PUT("/schedules/:id", handler.UpdateSchedule, middlewares.JWTMiddleware())

}

func (delivery *scheduleDelivery) PostSchedule(c echo.Context) error {
	var scheduleRequestdata ScheduleRequest

	errBind := c.Bind(&scheduleRequestdata)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	row, err := delivery.scheduleUsecase.PostData(ToCore(scheduleRequestdata))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
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

func (delivery *scheduleDelivery) GetScheduleDetail(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.scheduleUsecase.GetScheduleDetailById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCoreScheduleDetail(result)))

}

func (delivery *scheduleDelivery) DeleteSchedule(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)
	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	row, err := delivery.scheduleUsecase.DeleteSchedule(id_conv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail delete data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("rows affected 0, fail delete data"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success delete data"))
}

func (delivery *scheduleDelivery) UpdateSchedule(c echo.Context) error {

	var scheduleRequestdata ScheduleRequest

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)
	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	rowDel, errDel := delivery.scheduleUsecase.DeleteSchedule(id_conv)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail delete data"))
	}
	if rowDel != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("rows affected 0, fail delete data"))
	}

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

func (delivery *scheduleDelivery) UpdateScheduleDetail(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	var updateScheduleDetail ScheduleDetailRequest
	errBind := c.Bind(&updateScheduleDetail)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	photoUpdateCore := ToCoreScheduleDetail(updateScheduleDetail)
	photoUpdateCore.ID = uint(id_conv)

	row, err := delivery.scheduleUsecase.PutScheduleDetail(photoUpdateCore, id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}
