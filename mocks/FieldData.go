package moks

import (
	"capstone-project/features/field"

	"github.com/stretchr/testify/mock"
)

type FieldData struct {
	mock.Mock
}

func (_m *FieldData) SelectAllField(venue_id int) (data []field.FieldCore, err error) {
	ret := _m.Called(venue_id)

	var r0 []field.FieldCore
	if rf, ok := ret.Get(0).(func(int) []field.FieldCore); ok {
		r0 = rf(venue_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]field.FieldCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(venue_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *FieldData) SelectFieldById(id int) (data field.FieldCore, err error) {
	ret := _m.Called(id)

	var r0 field.FieldCore
	if rf, ok := ret.Get(0).(func(int) field.FieldCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(field.FieldCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *FieldData) InsertData(data field.FieldCore) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(field.FieldCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(field.FieldCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *FieldData) UpdateField(data field.FieldCore, field_id int) (row int, err error) {
	ret := _m.Called(data, field_id)

	var r0 int
	if rf, ok := ret.Get(0).(func(field.FieldCore, int) int); ok {
		r0 = rf(data, field_id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(field.FieldCore, int) error); ok {
		r1 = rf(data, field_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *FieldData) DeleteField(id int) (row int, err error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFieldata interface {
	mock.TestingT
	Cleanup(func())
}

func NewFIeldData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
