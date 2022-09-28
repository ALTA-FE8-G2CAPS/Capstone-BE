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
	e.POST("/users", handler.RegisterUser, middlewares.JWTMiddleware())
	e.POST("/login", handler.LoginUser)
	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetUserById, middlewares.JWTMiddleware())
}

func (handler *userDelivery) LoginUser(c echo.Context) error {
	data := user.UserCore{}
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("user doesn't exist"))
	}

	token, err := handler.userUsecase.PostLogin(data)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "login success",
		"token":   token,
	})
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
