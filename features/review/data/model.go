package data

import (
	"capstone-project/features/review"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Name_user   string
	UserID      uint
	Nama_venue  string
	VenueID     uint
	Rate        uint
	Feedback    string
	Foto_review string
	Venue       Venue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	Fields            []Field  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reviews           []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	Venues       []Venue  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reviews      []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func fromCore(data review.ReviewCore) Review {
	return Review{
		Name_user:   data.Name_user,
		UserID:      data.UserID,
		Nama_venue:  data.Nama_venue,
		VenueID:     data.VenueID,
		Rate:        data.Rate,
		Feedback:    data.Feedback,
		Foto_review: data.Foto_review,
	}
}

func (data *Review) toCore() review.ReviewCore {
	return review.ReviewCore{
		ID:          data.ID,
		Name_user:   data.Name_user,
		UserID:      data.UserID,
		Nama_venue:  data.Nama_venue,
		VenueID:     data.VenueID,
		Rate:        data.Rate,
		Feedback:    data.Feedback,
		Foto_review: data.Foto_review,
	}
}

func toCoreList(data []Review) []review.ReviewCore {
	var list []review.ReviewCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
