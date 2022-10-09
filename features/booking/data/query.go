package data

import (
	"capstone-project/features/booking"
	"errors"
	"fmt"

	"github.com/midtrans/midtrans-go/coreapi"
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
	fmt.Println("isinya", newBooking.VenueID)
	tx := repo.db.Create(&newBooking)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *bookingData) SelectAllBooking(user_id, field_id, venue_id int) ([]booking.BookingCore, error) {
	var dataBooking []Booking

	if user_id != 0 && field_id != 0 && venue_id != 0 {
		tx := repo.db.Where("user_id = ? AND field_id=? AND venue_id = ?", user_id, field_id, venue_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)
		// fmt.Println(dataField[0].Venue.Name_venue)
		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if user_id != 0 && field_id != 0 {
		tx := repo.db.Where("user_id = ? AND field_id=?", user_id, field_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)

		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if user_id != 0 && venue_id != 0 {
		tx := repo.db.Where("user_id = ? AND venue_id=?", user_id, venue_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)

		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if field_id != 0 && venue_id != 0 {
		tx := repo.db.Where("field_id = ? AND venue_id=?", field_id, venue_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)

		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if venue_id != 0 {
		tx := repo.db.Where("venue_id = ?", venue_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)

		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if user_id != 0 {
		tx := repo.db.Where("user_id = ?", user_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)

		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else if field_id != 0 {
		tx := repo.db.Where("field_id = ?", field_id).Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)
		// fmt.Println(dataField[0].Venue.Name_venue)
		if tx.Error != nil {
			return []booking.BookingCore{}, tx.Error
		}
	} else {

		tx := repo.db.Preload("User").Preload("Field.Venue").Preload("ScheduleDetail").Find(&dataBooking)
		// fmt.Println("ini adalaj", dataBooking[0].ScheduleDetail)

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

func (repo *bookingData) DeleteBooking(booking_id int) (int, error) {
	var dataBooking Booking

	tx := repo.db.Where("id = ?", booking_id).Unscoped().Delete(&dataBooking)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed, rows affected 0")
	}
	return int(tx.RowsAffected), nil
}

func (repo *bookingData) UpdatePayment(data booking.BookingCore, booking_id int) (int, error) {
	var bookingUpdate Booking
	txDataOld := repo.db.First(&bookingUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.FieldID != 0 {
		bookingUpdate.FieldID = data.FieldID
	}

	if data.UserID != 0 {
		bookingUpdate.UserID = data.UserID
	}

	if data.ScheduleDetailID != 0 {
		bookingUpdate.ScheduleDetailID = data.ScheduleDetailID
	}

	if data.Payment_method != "" {
		bookingUpdate.Payment_method = data.Payment_method
	}
	if data.OrderID != "" {
		bookingUpdate.OrderID = data.OrderID
	}
	if data.TransactionID != "" {
		bookingUpdate.TransactionID = data.TransactionID
	}
	if data.Virtual_account != "" {
		bookingUpdate.Virtual_account = data.Virtual_account
	}
	if data.Transaction_time != "" {
		bookingUpdate.Transaction_time = data.Transaction_time
	}
	if data.Transaction_time != "" {
		bookingUpdate.Transaction_exp = data.Transaction_exp
	}
	tx := repo.db.Save(&bookingUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *bookingData) CreateDataPayment(reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	payment, errPayment := coreapi.ChargeTransaction(&reqPay)
	if errPayment != nil {
		return nil, errPayment.RawError
	}
	return payment, nil
}

func (repo *bookingData) UpdatepaymentWebhook(data booking.BookingCore) (int, error) {
	tx := repo.db.Model(&Booking{}).Where("order_id = ?", data.OrderID).Select("status_payment").Update("status_payment", data.Status_payment)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return 0, nil
}
