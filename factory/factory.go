package factory

import (
	userData "capstone-project/features/user/data"
	userDelivery "capstone-project/features/user/delivery"
	userUsecase "capstone-project/features/user/usecase"

	venueData "capstone-project/features/venue/data"
	venueDelivery "capstone-project/features/venue/delivery"
	venueUsecase "capstone-project/features/venue/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	venueDataFactory := venueData.New(db)
	venueUsecaseFactory := venueUsecase.New(venueDataFactory)
	venueDelivery.New(e, venueUsecaseFactory)

}
