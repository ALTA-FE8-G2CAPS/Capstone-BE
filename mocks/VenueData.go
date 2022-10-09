package moks

import (
	"capstone-project/features/venue"

	"github.com/stretchr/testify/mock"
)

type VenueData struct {
	mock.Mock
}

func (_m *VenueData) SelectAllVenue(user_id int) (data []venue.VenueCore, err error) {
	ret := _m.Called(user_id)

	var r0 []venue.VenueCore
	if rf, ok := ret.Get(0).(func(int) []venue.VenueCore); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]venue.VenueCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VenueData) SelectVenueById(id int) (data venue.VenueCore, err error) {
	ret := _m.Called(id)

	var r0 venue.VenueCore
	if rf, ok := ret.Get(0).(func(int) venue.VenueCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(venue.VenueCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VenueData) InsertData(data venue.VenueCore) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(venue.VenueCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(venue.VenueCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VenueData) UpdateVenue(data venue.VenueCore, user_id int) (row int, err error) {
	ret := _m.Called(data, user_id)

	var r0 int
	if rf, ok := ret.Get(0).(func(venue.VenueCore, int) int); ok {
		r0 = rf(data, user_id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(venue.VenueCore, int) error); ok {
		r1 = rf(data, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VenueData) DeleteVenue(user_id, venue_id int) (row int, err error) {
	ret := _m.Called(user_id, venue_id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(user_id, venue_id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(user_id, venue_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VenueData) UploadPhoto(data venue.FotoVenue) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(venue.FotoVenue) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(venue.FotoVenue) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *VenueData) UpdatePhoto(data venue.FotoVenue, foto_venue_id int) (row int, err error) {
	ret := _m.Called(data, foto_venue_id)

	var r0 int
	if rf, ok := ret.Get(0).(func(venue.FotoVenue, int) int); ok {
		r0 = rf(data, foto_venue_id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(venue.FotoVenue, int) error); ok {
		r1 = rf(data, foto_venue_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewVenueData interface {
	mock.TestingT
	Cleanup(func())
}

func NewVenueData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
