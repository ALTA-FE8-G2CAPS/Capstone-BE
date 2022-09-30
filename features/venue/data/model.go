package data

import (
	"capstone-project/features/venue"

	"gorm.io/gorm"
)

type Venue struct {
	gorm.Model
	Name_venue        string
	Address_venue     string
	UserID            uint
	Description_venue string
	Latitude          float64
	Longitude         float64
	FotoVenues        []FotoVenue `gorm:"foreignKey:VenueID"`
	User              User
}

type FotoVenue struct {
	VenueID    uint `gorm:"primarykey"`
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
	Venues       []Venue
}

type Owner struct {
	UserID     uint
	Foto_owner string
}

func fromCore(dataCore venue.VenueCore) Venue {
	return Venue{
		Name_venue:        dataCore.Name_venue,
		Address_venue:     dataCore.Address_venue,
		UserID:            dataCore.UserID,
		Latitude:          dataCore.Latitude,
		Longitude:         dataCore.Longitude,
		Description_venue: dataCore.Description_venue,
	}
}

func (dataVenue *Venue) toCore() venue.VenueCore {
	return venue.VenueCore{
		ID:                dataVenue.ID,
		Name_venue:        dataVenue.Name_venue,
		Address_venue:     dataVenue.Address_venue,
		Description_venue: dataVenue.Description_venue,
		Name_user:         dataVenue.User.Name_User,
		Latitude:          dataVenue.Latitude,
		Longitude:         dataVenue.Longitude,
	}
}

func toCoreList(dataVenue []Venue) []venue.VenueCore {
	var dataCore []venue.VenueCore

	for key := range dataVenue {
		dataCore = append(dataCore, dataVenue[key].toCore())
	}
	return dataCore
}
