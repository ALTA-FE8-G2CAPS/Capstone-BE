package usecase

import (
	"capstone-project/features/review"
	"errors"
)

type reviewUsecase struct {
	reviewData review.DataInterface
}

func New(data review.DataInterface) review.Usecaseinterface {
	return &reviewUsecase{
		data,
	}
}

func (usecase *reviewUsecase) PostReview(data review.ReviewCore) (row int, err error) {
	if data.Rate == 0 || data.Feedback == "" {
		return 0, errors.New("data tidak boleh kosong")
	}
	row, err = usecase.reviewData.InsertReview(data)
	if err != nil {
		return 0, err
	}
	return row, err
}

func (usecase *reviewUsecase) GetReviewById(venue_id int) ([]review.ReviewCore, error) {
	result, err := usecase.reviewData.SelectReviewById(venue_id)
	if err != nil {
		return nil, errors.New("data tidak ditemukan")
	}
	return result, err
}
