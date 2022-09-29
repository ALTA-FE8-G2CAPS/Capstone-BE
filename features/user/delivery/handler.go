package delivery

import (
	"capstone-project/features/user"
	"capstone-project/middlewares"
	"capstone-project/utils/helper"
	"net/http"
	"strconv"

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
	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetUserById, middlewares.JWTMiddleware())
	// e.PUT("/users/:id", handler.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeleteUser, middlewares.JWTMiddleware())
}

func (handler *userDelivery) LoginUser(c echo.Context) error {
	var data UserRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("user doesn't exist"))
	}

	token, err := handler.userUsecase.PostLogin(ToCore(data))
	claim, _ := middlewares.ExtractClaims(token)
	role := claim["role"].(string)
	user := claim["name_user"].(string)
	user_owner := claim["user_owner"].(bool)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Fail_Resp("data doesn't exist"))
	}

	return c.JSON(http.StatusOK, helper.Success_Login("success login", token, role, user, user_owner))

}

func (handler *userDelivery) RegisterUser(c echo.Context) error {
	var data UserRequest
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to register"))
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

// func (handler *userDelivery) UpdateUser(c echo.Context) error {
// 	id := c.Param("id")
// 	idConv, errConv := strconv.Atoi(id)
// 	if errConv != nil {
// 		return c.JSON(400, map[string]interface{}{
// 			"message": errConv.Error(),
// 		})
// 	}
// 	var data UserRequest
// 	errBind := c.Bind(&data)

// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("fail to update"))
// 	}

// 	updateCore := ToCore(data)
// 	updateCore.ID = uint(idConv)

// 	row, err := handler.userUsecase.PutData(updateCore)
// 	if err != nil {
// 		return c.JSON(400, map[string]interface{}{
// 			"message": "update error",
// 		})
// 	}

// 	if row != 1 {
// 		return c.JSON(400, map[string]interface{}{"message": "failed to update"})
// 	}

// 	return c.JSON(200, map[string]interface{}{
// 		"message": "update success",
// 	})
// }

func (handler *userDelivery) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errConv.Error())
	}

	row, err := handler.userUsecase.DeleteUser(idConv)
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
