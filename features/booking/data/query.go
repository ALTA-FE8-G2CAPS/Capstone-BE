package data

import (
	"capstone-project/features/booking"

	"gorm.io/gorm"
)

type bookingData struct {
	db *gorm.DB
}

func New(db *gorm.DB) booking.DataInterface {
	return &bookingData{
		db: db,
	}

}

func (repo *bookingData) InsertData(data booking.BookingCore) (int, error) {
	newBooking := fromCore(data)
	// fmt.Println(newBooking.UserID)
	tx := repo.db.Create(&newBooking)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *bookingData) SelectAllBooking(user_id, field_id int) ([]booking.BookingCore, error) {
	var dataBooking []Booking

	if user_id != 0 && field_id != 0 {
		tx := repo.db.Where("user_id = ? AND field_id=?", user_id, field_id).Preload("User").Preload("Field.Venue").Find(&dataBooking)
		// fmt.Println(dataField[0].Venue.Name_venue)
		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}

	} else if user_id != 0 {
		tx := repo.db.Where("user_id = ?", user_id).Preload("User").Preload("Field.Venue").Find(&dataBooking)
		// fmt.Println(dataField[0].Venue.Name_venue)
		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if field_id != 0 {
		tx := repo.db.Where("field_id = ?", field_id).Preload("User").Preload("Field.Venue").Find(&dataBooking)
		// fmt.Println(dataField[0].Venue.Name_venue)
		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else {
		// tx := repo.db.Joins("User").Joins("Field").Joins("inner join venues on venues.id = fields.venue_id").Find(&dataBooking)
		tx := repo.db.Preload("User").Preload("Field.Venue").Find(&dataBooking)

		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	}
	return toCoreList(dataBooking), nil

}

func (repo *bookingData) SelectBookingById(id int) (booking.BookingCore, error) {
	var bookingData Booking
	bookingData.ID = uint(id)

	tx := repo.db.Where("id = ?", id).Preload("User").Preload("Field.Venue").Find(&bookingData)

	if tx.Error != nil {
		return booking.BookingCore{}, tx.Error
	}

	bookingDataCore := bookingData.toCore()
	return bookingDataCore, nil

}
