package delivery

import (
	"capstone-project/features/venue"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"fmt"
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
	userId := middlewares.ExtractToken(c)
	fmt.Println(userId)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}
	var venue_RequestData VenueRequest
	venue_RequestData.UserID = uint(userId)
	errBind := c.Bind(&venue_RequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	row, err := delivery.venueUsecase.PostData(ToCore(venue_RequestData))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}

	return c.JSON(http.StatusOK, helper.Success_Resp("success input"))

}
