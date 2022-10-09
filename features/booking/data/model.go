package data

import (
	"capstone-project/features/booking"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID           uint `gorm:"foreignKey:UserID"`
	FieldID          uint `gorm:"foreignKey:FieldID"`
	VenueID          uint `gorm:"foreignKey:VenueID"`
	ScheduleDetailID uint
	Total_price      uint
	Payment_method   string
	OrderID          string
	TransactionID    string
	Status_payment   string `gorm:"default:pending"`
	Virtual_account  string
	Transaction_time string
	Transaction_exp  string
	User             User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Field            Field          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ScheduleDetail   ScheduleDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Venue            Venue          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type Schedule struct {
	gorm.Model
	FieldID         uint `gorm:"foreignkey:FieldID"`
	Day             string
	Start_hours     string
	End_hours       string
	Field           Field
	ScheduleDetails []ScheduleDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type ScheduleDetail struct {
	gorm.Model
	ScheduleID      uint
	Start_hours     string
	End_hours       string
	Status_schedule string `gorm:"default:Available"`
}
type User struct {
	gorm.Model
	Name_User    string
	Email        string `gorm:"unique"`
	Password     string
	Role         string `gorm:"default:user"`
	Address_user string
	Foto_user    string
	User_owner   bool      `gorm:"default:false"`
	Owner        Owner     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Venues       []Venue   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bookings     []Booking `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Owner struct {
	UserID     uint `gorm:"primarykey"`
	Foto_owner string
}

type Venue struct {
	gorm.Model
	Name_venue        string
	Address_venue     string
	UserID            uint
	Description_venue string
	Latitude          float64
	Longitude         float64
	FotoVenues        []FotoVenue
	User              User
}

type FotoVenue struct {
	gorm.Model
	VenueID    uint
	Foto_venue string
}
type Field struct {
	gorm.Model
	VenueID   uint `gorm:"foreignKey:VenueID"`
	Category  string
	Price     uint
	Venue     Venue      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedules []Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func fromCore(data booking.BookingCore) Booking {
	return Booking{
		UserID:           data.UserID,
		FieldID:          data.FieldID,
		VenueID:          data.VenueID,
		ScheduleDetailID: data.ScheduleDetailID,
		Total_price:      data.Total_price,
		Payment_method:   data.Payment_method,
		TransactionID:    data.TransactionID,
		Status_payment:   data.Status_payment,
		Virtual_account:  data.Virtual_account,
		Transaction_time: data.Transaction_time,
		Transaction_exp:  data.Transaction_exp,
	}
}

func (data *Booking) toCore() booking.BookingCore {
	return booking.BookingCore{
		ID:               data.ID,
		UserID:           data.UserID,
		Name_User:        data.User.Name_User,
		Email:            data.User.Email,
		VenueID:          data.VenueID,
		Nama_venue:       data.Field.Venue.Name_venue,
		FieldID:          data.FieldID,
		Category:         data.Field.Category,
		ScheduleDetailID: data.ScheduleDetailID,
		Start_hours:      data.ScheduleDetail.Start_hours,
		End_Hours:        data.ScheduleDetail.End_hours,
		Price:            data.Field.Price,
		Total_price:      data.Total_price,
		Payment_method:   data.Payment_method,
		OrderID:          data.OrderID,
		TransactionID:    data.TransactionID,
		Status_payment:   data.Status_payment,
		Virtual_account:  data.Virtual_account,
		Transaction_time: data.Transaction_time,
		Transaction_exp:  data.Transaction_exp,
	}
}

func toCoreList(data []Booking) []booking.BookingCore {
	var list []booking.BookingCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
