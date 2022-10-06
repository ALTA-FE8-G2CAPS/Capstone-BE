package data

import (
	"capstone-project/features/booking/data"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID     uint
	Name_user  string
	VenueID    uint
	Nama_venue string
	TotalPrice uint
	Status     string
	Booking    data.Booking
}
