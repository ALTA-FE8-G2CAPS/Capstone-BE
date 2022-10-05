package delivery

import (
	"capstone-project/config"
	"capstone-project/features/user"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &userDelivery{
		userUsecase: usecase,
	}
	e.POST("/users", handler.RegisterUser)
	e.POST("/login", handler.LoginUser)
	e.POST("/users/owner", handler.RegisterOwner, middlewares.JWTMiddleware())
	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/request", handler.GetAllRequest, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users", handler.UpdateUser, middlewares.JWTMiddleware())
	e.PUT("/users/adminapprove/:id", handler.Approve, middlewares.JWTMiddleware())
	e.DELETE("/users", handler.DeleteUser, middlewares.JWTMiddleware())
}

func (handler *userDelivery) LoginUser(c echo.Context) error {
	var data UserRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("user doesn't exist"))
	}

	token, err := handler.userUsecase.PostLogin(ToCore(data))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("data doesn't exist"))
	}
	claim, _ := middlewares.ExtractClaims(token)
	role := claim["role"].(string)
	user := claim["name_user"].(string)
	user_owner := claim["user_owner"].(bool)
	user_id := claim["userId"].(float64)
	foto_user := claim["foto_user"].(string)

	return c.JSON(http.StatusOK, helper.Success_Login("success login", token, role, user, user_owner, user_id, foto_user))

}

func (handler *userDelivery) RegisterUser(c echo.Context) error {
	var data UserRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to register"))
	}
	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto_user")
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
		imageName := data.Name_User + "_" + "photo" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		data.Foto_user = imageaddress
	}
	row, err := handler.userUsecase.PostData(ToCore(data))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if row != 1 {
		return c.JSON(400, map[string]interface{}{"message": "failed to register"})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "register success",
		"row":     row,
	})
}

func (handler *userDelivery) GetAllUser(c echo.Context) error {
	data, err := handler.userUsecase.GetAllUser()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    FromCoreList(data),
	})
}

func (handler *userDelivery) GetUserById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(400, map[string]interface{}{
			"message": errConv.Error(),
		})
	}
	data, err := handler.userUsecase.GetUserById(idConv)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    FromCore(data),
	})
}

func (handler *userDelivery) UpdateUser(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}
	var data UserRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to update"))
	}
	updateCore := ToCore(data)
	updateCore.ID = uint(userId)

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto_user")
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
		imageName := data.Name_User + "_" + "photo" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		updateCore.Foto_user = imageaddress
	}

	row, err := handler.userUsecase.PutData(updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Fail Update User Data"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Update Row Affected Is Not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("Success Update Data"))
}

func (handler *userDelivery) DeleteUser(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}

	row, err := handler.userUsecase.DeleteUser(userId)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "delete error",
		})
	}

	if row != 1 {
		return c.JSON(400, map[string]interface{}{"message": "failed to delete"})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "delete success",
	})
}

func (handler *userDelivery) RegisterOwner(c echo.Context) error {
	userId := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail decrypt jwt token"))
	}

	var data OwnerRequest
	data.UserID = uint(userId)
	errBind := c.Bind(&data)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to create owner"))
	}
	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto_user")
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
		imageName := strconv.Itoa(int(data.UserID)) + "_" + "photo" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("fail to upload file"))
		}

		data.Foto_owner = imageaddress
	}

	row, err := handler.userUsecase.PostOwner(ToCoreOwner(data))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if row != 1 {
		return c.JSON(400, map[string]interface{}{"message": "failed to create owner"})
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("success create owner"))
}

func (handler *userDelivery) GetAllRequest(c echo.Context) error {
	data, err := handler.userUsecase.GetVerificationRequest()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    FromCoreList(data),
	})
}
func (handler *userDelivery) Approve(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(400, map[string]interface{}{
			"message": errConv.Error(),
		})
	}

	var data UserRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to update"))
	}

	updateCore := ToCore(data)
	updateCore.ID = uint(idConv)

	row, err := handler.userUsecase.AdminApprove(updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Failed approve request"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("Update Row Affected Is Not 1"))
	}
	return c.JSON(http.StatusOK, helper.Success_Resp("Success approve request "))
}
