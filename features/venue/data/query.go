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
