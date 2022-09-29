package data

import (
	"capstone-project/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama_User    string `json:"nama_user" form:"nama_user"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Role         string `json:"role" form:"role"`
	Address_user string `json:"address_user" form:"address_user"`
	Foto_user    string
	User_owner   string `json:"user_owner" form:"user_owner"`
	Foto_owner   Owner
	Venues       []Venue
}

type Owner struct {
	UserID     uint
	Foto_owner string
}

type Venue struct {
	gorm.Model
	Nama_venue        string
	Address_venue     string
	UserID            uint
	Description_venue string
	Latitude          string
	Longitude         string
	User              User
	Fields            []Field
	Reviews           []Review
	FotoVenues        []FotoVenue
}

type FotoVenue struct {
	gorm.Model
	Foto_venue string
}

type Field struct {
	gorm.Model
	VenueID    uint
	Category   string
	Price      int
	FotoFields []FotoField
}

type FotoField struct {
	gorm.Model
	Foto_field string
}

type Review struct {
	gorm.Model
	VenueID     uint
	UserID      uint
	Rate        int
	Feedback    string
	Foto_review string
}

func fromCore(data user.UserCore) User {
	return User{
		Nama_User:    data.Nama_User,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Address_user: data.Address_user,
		Foto_user:    data.Foto_user,
		User_owner:   data.User_owner,
	}
}

func (data *User) toCore() user.UserCore {
	return user.UserCore{
		Nama_User:    data.Nama_User,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Address_user: data.Address_user,
		Foto_user:    data.Foto_user,
		User_owner:   data.User_owner,
	}
}

func toCoreList(data []User) []user.UserCore {
	var list []user.UserCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
