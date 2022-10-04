package factory

import (
	userData "capstone-project/features/user/data"
	userDelivery "capstone-project/features/user/delivery"
	userUsecase "capstone-project/features/user/usecase"

	venueData "capstone-project/features/venue/data"
	venueDelivery "capstone-project/features/venue/delivery"
	venueUsecase "capstone-project/features/venue/usecase"

	fieldData "capstone-project/features/field/data"
	fieldDelivery "capstone-project/features/field/delivery"
	fieldUsecase "capstone-project/features/field/usecase"

	bookingData "capstone-project/features/booking/data"
	bookingDelivery "capstone-project/features/booking/delivery"
	bookingUsecase "capstone-project/features/booking/usecase"

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

	fieldDataFactory := fieldData.New(db)
	fieldUsecaseFactory := fieldUsecase.New(fieldDataFactory)
	fieldDelivery.New(e, fieldUsecaseFactory)

	bookingDataFactory := bookingData.New(db)
	bookingUsecaseFactory := bookingUsecase.New(bookingDataFactory, fieldDataFactory)
	bookingDelivery.New(e, bookingUsecaseFactory)

}
