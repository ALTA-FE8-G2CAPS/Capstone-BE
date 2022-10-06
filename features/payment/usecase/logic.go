package usecase

import (
	"capstone-project/features/payment"
	"errors"
)

type paymentUsecase struct {
	paymentData payment.DataInterface
}

func New(pymt payment.DataInterface) payment.UsecaseInterface {
	return &paymentUsecase{
		paymentData: pymt,
	}
}

func (usecase *paymentUsecase) CreatePayment(userId int) (int, int, error) {
	if userId == 0 {
		return -1, -1, errors.New("user id is empty")
	}

	bookingId, totalPrice, err := usecase.paymentData.DataPayment(userId)
	if err != nil {
		return -1, -1, errors.New("fail operation")
	}
	return bookingId, totalPrice, nil
}
