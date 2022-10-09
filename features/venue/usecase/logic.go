package usecase

import (
	"capstone-project/features/venue"
	"errors"
	"math"
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
	if data.Name_venue == "" || data.Address_venue == "" || data.Description_venue == "" || data.Latitude == 0 || data.Longitude == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}
	row, err = usecase.venueData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *venueUsecase) GetAllVenue(user_id int) ([]venue.VenueCore, error) {
	dataMentee, err := usecase.venueData.SelectAllVenue(user_id)

	for index, mentee := range dataMentee {
		max := -1
		min := math.MaxInt64
		for _, price := range mentee.Price {
			if int(price.Price) >= max {
				max = int(price.Price)
			}
			if int(price.Price) <= min {
				min = int(price.Price)
			}
		}
		if len(mentee.Price) > 0 {
			dataMentee[index].Max_price = uint(max)
			dataMentee[index].Min_price = uint(min)
		}

	}

	// fmt.Println(dataMentee[0].Max_price)
	return dataMentee, err

}

func (usecase *venueUsecase) GetVenueById(id int) (venue.VenueCore, error) {
	result, err := usecase.venueData.SelectVenueById(id)
	if err != nil {
		return venue.VenueCore{}, err
	}
	return result, nil
}

func (usecase *venueUsecase) PutData(data venue.VenueCore, user_id int) (int, error) {
	row, err := usecase.venueData.UpdateVenue(data, user_id)
	return row, err
}

func (usecase *venueUsecase) DeleteVenue(user_id, venue_id int) (int, error) {
	result, err := usecase.venueData.DeleteVenue(user_id, venue_id)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (usecase *venueUsecase) PostPhoto(data venue.FotoVenue) (row int, err error) {

	row, err = usecase.venueData.UploadPhoto(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *venueUsecase) PutPhoto(data venue.FotoVenue, foto_venue_id int) (int, error) {
	row, err := usecase.venueData.UpdatePhoto(data, foto_venue_id)
	if err != nil {
		return -1, err
	}
	return row, nil
}
