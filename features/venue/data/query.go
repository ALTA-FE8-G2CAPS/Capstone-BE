package data

import (
	"capstone-project/features/venue"
	"errors"

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

func (repo *venueData) SelectVenueById(id int) (venue.VenueCore, error) {
	var dataVenue Venue
	dataVenue.ID = uint(id)

	tx := repo.db.Where("id = ?", id).Preload("User").First(&dataVenue)

	if tx.Error != nil {
		return venue.VenueCore{}, tx.Error
	}

	dataVenueCore := dataVenue.toCore()
	return dataVenueCore, nil

}

func (repo *venueData) DeleteVenue(user_id, venue_id int) (row int, err error) {
	var dataVenue Venue
	dataVenue.ID = uint(venue_id)

	tx := repo.db.Where("user_id = ?", user_id).Unscoped().Delete(&dataVenue)

	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *venueData) UpdateVenue(data venue.VenueCore, user_id int) (int, error) {
	var venueUpdate Venue
	txDataOld := repo.db.Where("user_id = ?", user_id).First(&venueUpdate, data.ID)
	// result := repo.db.Model(&Mentee{}).Where("id = ?", data.ID).Updates(fromCore(data))
	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Name_venue != "" {
		venueUpdate.Name_venue = data.Name_venue
	}

	if data.Address_venue != "" {
		venueUpdate.Address_venue = data.Address_venue
	}

	if data.Description_venue != "" {
		venueUpdate.Description_venue = data.Description_venue
	}

	tx := repo.db.Where("user_id = ?", user_id).Save(&venueUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected != 1 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx.RowsAffected), nil
}
