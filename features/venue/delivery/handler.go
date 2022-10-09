package delivery

import (
	"capstone-project/config"
	"capstone-project/features/venue"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"fmt"
	"net/http"
	"strconv"

	"time"

	"github.com/labstack/echo/v4"
)

type venueDelivery struct {
	venueUsecase venue.UsecaseInterface
}

func New(e *echo.Echo, usecase venue.UsecaseInterface) {
	handler := &venueDelivery{
		venueUsecase: usecase,
	}

	e.POST("/venues", handler.PostVenue, middlewares.JWTMiddleware())
	e.GET("/venues", handler.GetVenue, middlewares.JWTMiddleware())
	e.GET("/venues/:id", handler.GetVenueId, middlewares.JWTMiddleware())
	e.DELETE("/venues/:id", handler.DeleteVenue, middlewares.JWTMiddleware())
	e.PUT("/venues/:id", handler.UpdateVenue, middlewares.JWTMiddleware())
	e.POST("venues/foto/:id", handler.PostPhoto, middlewares.JWTMiddleware())
	e.PUT("venues/foto/:id", handler.UpdatePhoto, middlewares.JWTMiddleware())

}

func (delivery *venueDelivery) PostVenue(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	// fmt.Println(userId)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}
	var venue_RequestData VenueRequest
	venue_RequestData.UserID = uint(userId)
	fmt.Println(venue_RequestData.Name_venue)
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

func (delivery *venueDelivery) GetVenue(c echo.Context) error {

	user_id, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil && user_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail converse class_id param"))
	}

	dataMentee, err := delivery.venueUsecase.GetAllVenue(user_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("get all data success", FromCoreList(dataMentee)))
}

func (delivery *venueDelivery) GetVenueId(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.venueUsecase.GetVenueById(id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get data"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCore(result)))

}

func (delivery *venueDelivery) DeleteVenue(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)
	if user_id == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail operation"))
	}

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	row, err := delivery.venueUsecase.DeleteVenue(user_id, id_conv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to delete data (mungkin anda tidak mempunyai venue)"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to delete venue"))
	}
	return c.JSON(http.StatusOK, helper.Success_DataResp("success delete venue", row))
}

func (delivery *venueDelivery) UpdateVenue(c echo.Context) error {
	user_id := middlewares.ExtractToken(c)
	if user_id == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail operation"))
	}

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	var venueUpdate VenueRequest
	errBind := c.Bind(&venueUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	venueUpdateCore := ToCore(venueUpdate)
	venueUpdateCore.ID = uint(id_conv)

	row, err := delivery.venueUsecase.PutData(venueUpdateCore, user_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update data (mungkin anda tidak mempunyai venue)"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}

func (delivery *venueDelivery) PostPhoto(c echo.Context) error {
	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	var data Foto_venueRequest
	data.VenueID = uint(id_conv)

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto_venue")
	if fotoerr != http.ErrMissingFile || fotoerr == nil {
		format, errf := helper.CheckFile(infoFoto.Filename)
		if errf != nil {
			return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Format Error"))
		}
		//checksize
		err_image_size := helper.CheckSize(infoFoto.Size)
		if err_image_size != nil {
			return c.JSON(http.StatusBadRequest, err_image_size)
		}
		//rename
		waktu := fmt.Sprintf("%v", time.Now())
		imageName := strconv.Itoa(int(data.VenueID)) + "_" + "photo" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		data.Foto_venue = imageaddress
	}
	errBind := c.Bind(&data)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to upload photo"))
	}
	row, err := delivery.venueUsecase.PostPhoto(ToCoreFoto_venue(data))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if row != 1 {
		return c.JSON(400, map[string]interface{}{"message": "failed to upload photo"})
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success upload photo"))
}

func (delivery *venueDelivery) UpdatePhoto(c echo.Context) error {

	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)

	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}
	var photoUpdate Foto_venueRequest
	errBind := c.Bind(&photoUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind user data"))
	}

	photoUpdateCore := ToCoreFoto_venue(photoUpdate)
	photoUpdateCore.ID = uint(id_conv)

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto_venue")
	if fotoerr != http.ErrMissingFile || fotoerr == nil {
		format, errf := helper.CheckFile(infoFoto.Filename)
		if errf != nil {
			return c.JSON(http.StatusBadRequest, helper.Fail_Resp("Format Error"))
		}
		//checksize
		err_image_size := helper.CheckSize(infoFoto.Size)
		if err_image_size != nil {
			return c.JSON(http.StatusBadRequest, err_image_size)
		}
		//rename
		waktu := fmt.Sprintf("%v", time.Now())
		imageName := strconv.Itoa(int(photoUpdate.VenueID)) + "_" + "photo" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		photoUpdateCore.Foto_Venue = imageaddress

	}
	row, err := delivery.venueUsecase.PutPhoto(photoUpdateCore, id_conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail update data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("update row affected is not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success update data"))
}
