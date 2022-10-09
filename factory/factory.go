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

	scheduleData "capstone-project/features/schedule/data"
	scheduleDelivery "capstone-project/features/schedule/delivery"
	scheduleUsecase "capstone-project/features/schedule/usecase"

	reviewData "capstone-project/features/review/data"
	reviewDelivery "capstone-project/features/review/delivery"
	reviewUsecase "capstone-project/features/review/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.NewUserUsecase(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	venueDataFactory := venueData.New(db)
	venueUsecaseFactory := venueUsecase.New(venueDataFactory)
	venueDelivery.New(e, venueUsecaseFactory)

	fieldDataFactory := fieldData.New(db)
	fieldUsecaseFactory := fieldUsecase.New(fieldDataFactory)
	fieldDelivery.New(e, fieldUsecaseFactory)

	scheduleDataFactory := scheduleData.New(db)
	scheduleUsecaseFactory := scheduleUsecase.New(scheduleDataFactory)
	scheduleDelivery.New(e, scheduleUsecaseFactory)

	bookingDataFactory := bookingData.New(db)
	bookingUsecaseFactory := bookingUsecase.New(bookingDataFactory, fieldDataFactory, scheduleDataFactory)
	bookingDelivery.New(e, bookingUsecaseFactory)

	reviewDataFactory := reviewData.New(db)
	reviewUsecaseFactory := reviewUsecase.New(reviewDataFactory)
	reviewDelivery.New(e, reviewUsecaseFactory)
}
