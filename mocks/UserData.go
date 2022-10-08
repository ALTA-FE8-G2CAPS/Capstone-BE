package moks

import (
	"capstone-project/features/user"

	"github.com/stretchr/testify/mock"
)

type UserData struct {
	mock.Mock
}

func (_m *UserData) InsertData(data user.UserCore) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.UserCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) SelectAllUser() (data []user.UserCore, err error) {
	ret := _m.Called()

	var r0 []user.UserCore
	if rf, ok := ret.Get(0).(func() []user.UserCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.UserCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) LoginUser(data user.UserCore) (token string, err error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(user.UserCore) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) UpdateUser(data user.UserCore) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.UserCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) DeleteUser(id int) (row int, err error) {
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

func (_m *UserData) SelectUserById(id int) (data user.UserCore, err error) {
	ret := _m.Called(id)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(int) user.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) InsertOwner(data user.Owner) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.Owner) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.Owner) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) SelectVerificationRequest() (data []user.UserCore, err error) {
	ret := _m.Called()

	var r0 []user.UserCore
	if rf, ok := ret.Get(0).(func() []user.UserCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.UserCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserData) AdminApprove(data user.UserCore) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.UserCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
