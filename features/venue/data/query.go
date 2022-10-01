package data

import (
	"capstone-project/features/venue"

	"gorm.io/gorm"
)

type venueData struct {
	db *gorm.DB
}

func New(db *gorm.DB) venue.DataInterface {
	return &venueData{
		db: db,
	}

}

func (repo *venueData) InsertData(data venue.VenueCore) (int, error) {
	newVenue := fromCore(data)

	tx := repo.db.Create(&newVenue)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *venueData) SelectAllVenue(user_id int) ([]venue.VenueCore, error) {
	var dataVenue []Venue

	if user_id != 0 {
		tx := repo.db.Where("user_id = ?", user_id).Preload("User").Find(&dataVenue)

		if tx.Error != nil {
			return []venue.VenueCore{}, tx.Error
		}
	} else {
		tx := repo.db.Preload("User").Find(&dataVenue)

		if tx.Error != nil {
			return []venue.VenueCore{}, tx.Error
		}
	}

	return toCoreList(dataVenue), nil

}
