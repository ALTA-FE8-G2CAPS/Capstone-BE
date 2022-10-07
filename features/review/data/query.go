package data

import (
	"capstone-project/features/review"

	"gorm.io/gorm"
)

type reviewData struct {
	db *gorm.DB
}

func New(db *gorm.DB) review.DataInterface {
	return &reviewData{
		db: db,
	}
}

func (repo *reviewData) InsertReview(data review.ReviewCore) (int, error) {
	newReview := fromCore(data)

	tx := repo.db.Create(&newReview)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *reviewData) SelectReviewById(venue_id int) ([]review.ReviewCore, error) {
	var dataReview []Review

	tx := repo.db.Where("venue_id = ?", venue_id).Preload("User").Preload("Venue").Find(&dataReview)

	if tx.Error != nil {
		return []review.ReviewCore{}, tx.Error
	}

	return toCoreList(dataReview), nil
}
