package data

import (
	"capstone-project/features/booking"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID           uint `gorm:"foreignKey:UserID"`
	FieldID          uint `gorm:"foreignKey:UserID"`
	Start_hours      uint
	End_hours        uint
	Total_price      uint
	Payment_method   string
	OrderID          uint
	TransactionID    uint
	Status_payment   string `gorm:"default:pending"`
	Virtual_account  string
	Transaction_time string
	User             User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Field            Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	VenueID  uint `gorm:"foreignKey:VenueID"`
	Category string
	Price    uint
	Venue    Venue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func fromCore(data booking.BookingCore) Booking {
	return Booking{
		UserID:           data.UserID,
		FieldID:          data.FieldID,
		Start_hours:      data.Start_hours,
		End_hours:        data.End_hours,
		Total_price:      data.Total_price,
		Payment_method:   data.Payment_method,
		TransactionID:    data.TransactionID,
		Status_payment:   data.Status_payment,
		Virtual_account:  data.Virtual_account,
		Transaction_time: data.Transaction_time,
	}
}

func (data *Booking) toCore() booking.BookingCore {
	return booking.BookingCore{
		ID:               data.ID,
		UserID:           data.UserID,
		Name_User:        data.User.Name_User,
		Nama_venue:       data.Field.Venue.Name_venue,
		FieldID:          data.FieldID,
		Category:         data.Field.Category,
		Start_hours:      data.Start_hours,
		End_hours:        data.End_hours,
		Total_price:      data.Total_price,
		Payment_method:   data.Payment_method,
		OrderID:          data.OrderID,
		TransactionID:    data.TransactionID,
		Status_payment:   data.Status_payment,
		Virtual_account:  data.Virtual_account,
		Transaction_time: data.Transaction_time,
	}
}

func toCoreList(data []Booking) []booking.BookingCore {
	var list []booking.BookingCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
