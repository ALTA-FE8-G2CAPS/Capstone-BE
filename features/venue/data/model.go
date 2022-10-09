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
	FotoVenues        []FotoVenue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Fields            []Field     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User              User
}

type FotoVenue struct {
	gorm.Model
	VenueID    uint `gorm:"foreignKey:VenueID"`
	Foto_venue string
}
type Field struct {
	gorm.Model
	VenueID  uint `gorm:"foreignKey:VenueID"`
	Category string
	Price    uint
	Venue    Venue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func fromCoreFoto(data venue.FotoVenue) FotoVenue {
	return FotoVenue{
		VenueID:    data.VenueID,
		Foto_venue: data.Foto_Venue,
	}
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
		UserID:            dataVenue.UserID,
		Name_user:         dataVenue.User.Name_User,
		Latitude:          dataVenue.Latitude,
		Longitude:         dataVenue.Longitude,
		Foto_venue:        toCoreFotoList(dataVenue.FotoVenues),
		Price:             toCoreFieldList(dataVenue.Fields),
	}
}

func toCoreList(dataVenue []Venue) []venue.VenueCore {
	var dataCore []venue.VenueCore

	for key := range dataVenue {
		dataCore = append(dataCore, dataVenue[key].toCore())
	}
	return dataCore
}

func toCoreFotoList(dataVenue []FotoVenue) []venue.FotoVenue {
	var dataCore []venue.FotoVenue

	for key := range dataVenue {
		var foto_venue venue.FotoVenue
		foto_venue.ID = dataVenue[key].ID
		foto_venue.VenueID = dataVenue[key].VenueID
		foto_venue.Foto_Venue = dataVenue[key].Foto_venue
		dataCore = append(dataCore, foto_venue)
	}
	return dataCore
}

func toCoreFieldList(dataVenue []Field) []venue.Field2 {
	var dataCore []venue.Field2

	for key := range dataVenue {
		var field venue.Field2
		field.VenueID = dataVenue[key].VenueID
		field.Price = dataVenue[key].Price
		field.Category = dataVenue[key].Category
		dataCore = append(dataCore, field)
	}
	// fmt.Println(dataVenue)
	return dataCore
}
