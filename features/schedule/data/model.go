package data

import (
	"capstone-project/features/schedule"

	"gorm.io/gorm"
)

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
	ScheduleID      uint `gorm:"foreignkey:ScheduleID"`
	Start_hours     uint
	End_hours       uint
	Status_schedule string
}
type Field struct {
	gorm.Model
	VenueID   uint `gorm:"foreignKey:VenueID"`
	Category  string
	Price     uint
	Venue     Venue      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedules []Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func fromCore(data schedule.ScheduleCore) Schedule {
	return Schedule{
		FieldID:     data.FieldID,
		Day:         data.Day,
		Start_hours: data.Start_hours,
		End_hours:   data.End_hours,
	}
}

func (data *Schedule) toCore() schedule.ScheduleCore {
	return schedule.ScheduleCore{
		ID:          data.ID,
		FieldID:     data.FieldID,
		Category:    data.Field.Category,
		Day:         data.Day,
		Start_hours: data.Start_hours,
		End_hours:   data.End_hours,
	}
}

func toCoreList(data []Schedule) []schedule.ScheduleCore {
	var list []schedule.ScheduleCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}