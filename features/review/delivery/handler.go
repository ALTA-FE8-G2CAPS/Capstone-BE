package delivery

import (
	"capstone-project/config"
	"capstone-project/features/review"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type reviewDelivery struct {
	reviewUsecase review.Usecaseinterface
}

func New(e *echo.Echo, usecase review.Usecaseinterface) {
	handler := &reviewDelivery{
		reviewUsecase: usecase,
	}

	e.POST("/reviews", handler.PostReview, middlewares.JWTMiddleware())
	e.GET("/reviews/:id", handler.GetReviewById, middlewares.JWTMiddleware())
}

func (delivery *reviewDelivery) PostReview(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail extract token"))
	}

	var reviewRequest ReviewRequest
	reviewRequest.UserID = uint(userId)
	errBind := c.Bind(&reviewRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail bind data"))
	}

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto_review")
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
		imageName := strconv.Itoa(int(reviewRequest.UserID)) + "_" + strconv.Itoa(int(reviewRequest.VenueID)) + "photo" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		reviewRequest.Foto_review = imageaddress
	}

	row, err := delivery.reviewUsecase.PostReview(ToCore(reviewRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail input data"))
	}
	fmt.Println(err)
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to input data"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success post review"))
}

func (delivery *reviewDelivery) GetReviewById(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail extract token"))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail convert param"))
	}

	review, err := delivery.reviewUsecase.GetReviewById(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail get review"))
	}

	return c.JSON(http.StatusOK, helper.Success_DataResp("success get data", FromCoreList(review)))
}
