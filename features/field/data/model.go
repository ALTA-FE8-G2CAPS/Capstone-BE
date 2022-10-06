package data

import (
	"capstone-project/features/field"

	"gorm.io/gorm"
)

type Field struct {
	gorm.Model
	VenueID   uint `gorm:"foreignKey:VenueID"`
	Category  string
	Price     uint
	Venue     Venue      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedules []Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Schedule struct {
	gorm.Model
	FieldID        uint `gorm:"foreignKey:FieldID"`
	Days           string
	Start_hours    uint
	End_hours      uint
	ScheduleDetail []ScheduleDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ScheduleDetail struct {
	gorm.Model
	ScheduleID      uint `gorm:"foreignkey:ScheduleID"`
	Start_hours     string
	End_hours       string
	Status_schedule string
}

type Venue struct {
	gorm.Model
	Name_venue        string
	Address_venue     string
	UserID            uint
	Description_venue string
	Latitude          float64
	Longitude         float64
	FotoVenues        []FotoVenue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User              User
	Fields            []Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type FotoVenue struct {
	gorm.Model
	VenueID    uint `gorm:"foreignKey:VenueID"`
	Foto_venue string
}

type User struct {
	gorm.Model
	Name_User    string
	Email        string `gorm:"unique"`
	Password     string
	Role         string `gorm:"default:user"`
	Address_user string
	Foto_user    string
	User_owner   bool `gorm:"default:false"`
	Owner        Owner
	Venues       []Venue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Owner struct {
	UserID     uint
	Foto_owner string
}

func fromCore(data field.FieldCore) Field {
	return Field{
		VenueID:  data.VenueID,
		Category: data.Category,
		Price:    data.Price,
	}
}

func (data *Field) toCore() field.FieldCore {
	return field.FieldCore{
		ID:         data.ID,
		VenueID:    data.VenueID,
		Name_venue: data.Venue.Name_venue,
		Category:   data.Category,
		Price:      data.Price,
	}
}

func toCoreList(data []Field) []field.FieldCore {
	var list []field.FieldCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
