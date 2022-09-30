package usecase

import (
	"capstone-project/features/venue"
	"errors"
)

type venueUsecase struct {
	venueData venue.DataInterface
}

func New(data venue.DataInterface) venue.UsecaseInterface {
	return &venueUsecase{
		data,
	}
}

func (usecase *venueUsecase) PostData(data venue.VenueCore) (row int, err error) {
	if data.Name_Venue == "" || data.Addres_Venue == "" || data.Description == "" || data.Latitude == 0 || data.Longitude == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}
	row, err = usecase.venueData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, err
}
