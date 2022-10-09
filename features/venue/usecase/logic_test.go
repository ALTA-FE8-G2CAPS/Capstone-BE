package usecase

import (
	"capstone-project/features/venue"
	mocks "capstone-project/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllVenue(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := []venue.VenueCore{{ID: 1, Name_venue: "gor bagus coba", Address_venue: "jalan jalan", Description_venue: "bagus", UserID: 1, Name_user: "Jono", Latitude: 1.1, Longitude: 1.1, Foto_venue: []venue.FotoVenue{{ID: 1, VenueID: 1, Foto_Venue: "foto.jpg"}}}}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("SelectAllData").Return(returnData, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetAllVenue(1)
		assert.NoError(t, err)
		assert.Equal(t, resultData[0].ID, returnData[0].ID)
		repo.AssertExpectations(t)
	})
}

func TestGetVenueById(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := venue.VenueCore{
		ID:                1,
		Name_venue:        "gor bagus coba",
		Address_venue:     "jalan jalan",
		Description_venue: "bagus",
		UserID:            1,
		Name_user:         "Jono",
		Latitude:          1.1,
		Longitude:         1.1,
		Foto_venue:        []venue.FotoVenue{{ID: 1, VenueID: 1, Foto_Venue: "foto.jpg"}},
	}
	t.Run("Success Get Data By Id", func(t *testing.T) {
		repo.On("SelectDataById", 1).Return(returnData, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetVenueById(1)
		assert.NoError(t, err)
		assert.Equal(t, resultData.ID, returnData.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Data By Id", func(t *testing.T) {
		repo.On("SelectDataById", 0).Return(venue.VenueCore{}, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetVenueById(0)
		assert.Error(t, err)
		assert.Equal(t, resultData.ID, 0)
		repo.AssertExpectations(t)
	})
}

func TestPostData(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := venue.VenueCore{
		ID:                1,
		Name_venue:        "gor bagus coba",
		Address_venue:     "jalan jalan",
		Description_venue: "bagus",
		UserID:            1,
		Name_user:         "Jono",
		Latitude:          1.1,
		Longitude:         1.1,
		Foto_venue:        []venue.FotoVenue{{ID: 1, VenueID: 1, Foto_Venue: "foto.jpg"}},
	}
	t.Run("Success Post Data", func(t *testing.T) {
		repo.On("InsertData", returnData).Return(1, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.PostData(returnData)
		assert.NoError(t, err)
		assert.Equal(t, resultData, 1)
		repo.AssertExpectations(t)
	})
}

func TestPutData(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := venue.VenueCore{
		ID:                1,
		Name_venue:        "gor bagus coba",
		Address_venue:     "jalan jalan",
		Description_venue: "bagus",
		UserID:            1,
		Name_user:         "Jono",
		Latitude:          1.1,
		Longitude:         1.1,
		Foto_venue:        []venue.FotoVenue{{ID: 1, VenueID: 1, Foto_Venue: "foto.jpg"}},
	}
	t.Run("Success Put Data", func(t *testing.T) {
		repo.On("UpdateData", returnData).Return(nil).Once()

		usecase := New(repo)
		_, err := usecase.PutData(returnData, int(returnData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := venue.VenueCore{
		ID:                1,
		Name_venue:        "gor bagus coba",
		Address_venue:     "jalan jalan",
		Description_venue: "bagus",
		UserID:            1,
		Name_user:         "Jono",
		Latitude:          1.1,
		Longitude:         1.1,
		Foto_venue:        []venue.FotoVenue{{ID: 1, VenueID: 1, Foto_Venue: "foto.jpg"}},
	}
	t.Run("Success Delete Data", func(t *testing.T) {
		repo.On("DeleteData", returnData).Return(nil).Once()

		usecase := New(repo)
		_, err := usecase.DeleteVenue(int(returnData.ID), int(returnData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestPostPhoto(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := venue.FotoVenue{
		ID:         1,
		VenueID:    1,
		Foto_Venue: "foto.jpg",
	}
	t.Run("Success Post Photo", func(t *testing.T) {
		repo.On("InsertPhoto", returnData).Return(1, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.PostPhoto(returnData)
		assert.NoError(t, err)
		assert.Equal(t, resultData, 1)
		repo.AssertExpectations(t)
	})
}

func TestUpdatePhoto(t *testing.T) {
	repo := new(mocks.VenueData)
	returnData := venue.FotoVenue{
		ID:         1,
		VenueID:    1,
		Foto_Venue: "foto.jpg",
	}
	t.Run("Success Update Photo", func(t *testing.T) {
		repo.On("UpdatePhoto", returnData).Return(nil).Once()

		usecase := New(repo)
		_, err := usecase.PutPhoto(returnData, int(returnData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
