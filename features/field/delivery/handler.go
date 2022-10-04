package delivery

import (
	"capstone-project/features/field"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type fieldDeliveryvery struct {
	fieldUsecase field.UsecaseInterface
}

func New(e *echo.Echo, usecase field.UsecaseInterface) {
	handler := &fieldDeliveryvery{
		fieldUsecase: usecase,
	}

	e.POST("/fields", handler.PostField, middlewares.JWTMiddleware())
	e.GET("/fields", handler.GetField, middlewares.JWTMiddleware())
	e.GET("/fields/:id", handler.GetFieldId, middlewares.JWTMiddleware())

}

func (delivery *fieldDeliveryvery) PostField(c echo.Context) error {

	var fieldRequestdata FieldRequest
	errBind := c.Bind(&fieldRequestdata)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	row, err := delivery.fieldUsecase.PostData(ToCore(fieldRequestdata))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input"))

}

func (delivery *fieldDeliveryvery) GetField(c echo.Context) error {

	venue_id, err := strconv.Atoi(c.QueryParam("venue_id"))
	if err != nil && venue_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail converse class_id param"))
	}

	data, err := delivery.fieldUsecase.GetAllField(venue_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all data success", FromCoreList(data)))
}

func (delivery *fieldDeliveryvery) GetFieldId(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.fieldUsecase.GetFieldById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}
