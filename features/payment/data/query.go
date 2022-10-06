package data

import (
	"capstone-project/features/payment"

	"gorm.io/gorm"
)

type paymentData struct {
	db *gorm.DB
}

func New(conn *gorm.DB) payment.DataInterface {
	return &paymentData{
		db: conn,
	}
}

func (repo *paymentData) DataPayment(userId int) (int, int, error) {
	var bookingId int
	var totalPrice int
	booking := repo.db.Raw("SELECT id FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", userId).Scan(&bookingId)
	if booking.Error != nil {
		return -1, -1, booking.Error
	}

	resTotalPrice := repo.db.Raw("SELECT total_price FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", userId).Scan(&totalPrice)
	if resTotalPrice.Error != nil {
		return -1, -1, resTotalPrice.Error
	}
	return bookingId, totalPrice, nil
}
