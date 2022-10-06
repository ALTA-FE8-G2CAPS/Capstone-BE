package migration

import (
	bookingModel "capstone-project/features/booking/data"
	fieldModel "capstone-project/features/field/data"
	scheduleModel "capstone-project/features/schedule/data"
	userModel "capstone-project/features/user/data"
	venueModel "capstone-project/features/venue/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&userModel.Owner{})
	db.AutoMigrate(&venueModel.Venue{})
	db.AutoMigrate(&venueModel.FotoVenue{})
	db.AutoMigrate(&fieldModel.Field{})
	db.AutoMigrate(&bookingModel.Booking{})
	db.AutoMigrate(&bookingModel.Booking{})
	db.AutoMigrate(&scheduleModel.Schedule{})
	db.AutoMigrate(&scheduleModel.ScheduleDetail{})

}
