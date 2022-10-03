package data

import (
	"capstone-project/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name_User    string
	Email        string `gorm:"unique"`
	Password     string
	Role         string `gorm:"default:user"`
	Address_user string
	Foto_user    string
	User_owner   bool    `gorm:"default:false"`
	Owner        Owner   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Venues       []Venue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

// type Field struct {
// 	gorm.Model
// 	VenueID    uint
// 	Category   string
// 	Price      int
// 	FotoFields []FotoField
// }

// type FotoField struct {
// 	gorm.Model
// 	Foto_field string
// }

// type Review struct {
// 	gorm.Model
// 	VenueID     uint
// 	UserID      uint
// 	Rate        int
// 	Feedback    string
// 	Foto_review string
// }

func fromCore(data user.UserCore) User {
	return User{
		Name_User:    data.Name_User,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Address_user: data.Address_user,
		Foto_user:    data.Foto_user,
		User_owner:   data.User_owner,
	}
}

func fromCoreOwner(data user.Owner) Owner {
	return Owner{
		UserID:     data.UserID,
		Foto_owner: data.Foto_owner,
	}
}

func (data *User) toCore() user.UserCore {
	return user.UserCore{
		ID:           data.ID,
		Name_User:    data.Name_User,
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
