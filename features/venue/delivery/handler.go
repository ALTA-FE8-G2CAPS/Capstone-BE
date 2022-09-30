package delivery

import (
	"capstone-project/features/venue"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VenueDelivery struct {
	venueUsecase venue.UsecaseInterface
}

func New(e *echo.Echo, usecase venue.UsecaseInterface) {
	handler := &VenueDelivery{
		venueUsecase: usecase,
	}

	e.POST("/venues", handler.PostVenue, middlewares.JWTMiddleware())

}

func (delivery *VenueDelivery) PostVenue(c echo.Context) error {
	var venue_RequestData VenueRequest
	errBind := c.Bind(&venue_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind produk data"))
	}

	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}
	venue_RequestData.UserID = uint(userId)

	row, err := delivery.venueUsecase.PostData(ToCore(venue_RequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input produk data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input produk data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input new produk"))

}
