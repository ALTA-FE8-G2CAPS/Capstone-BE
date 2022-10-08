package usecase

import (
	"capstone-project/features/user"
	mocks "capstone-project/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUser(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []user.UserCore{{ID: 1, Name_User: "Jono", Email: "example@gmail.com", Password: "qwerty", Address_user: "konohagakure", Role: "user", Foto_user: "foto.jpg", User_owner: false}}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("SelectAllData").Return(returnData, nil).Once()

		usecase := NewUserUsecase(repo)
		resultData, err := usecase.GetAllUser()
		assert.NoError(t, err)
		assert.Equal(t, resultData[0].ID, returnData[0].ID)
		repo.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := user.UserCore{
		ID:           1,
		Name_User:    "Jono",
		Email:        "example@gmail.com",
		Password:     "qwerty",
		Role:         "user",
		Address_user: "konohagakure",
		Foto_user:    "foto.jpg",
		User_owner:   false,
	}
	t.Run("Success Get Data By Id", func(t *testing.T) {
		repo.On("SelectDataById", 1).Return(returnData, nil).Once()

		usecase := NewUserUsecase(repo)
		resultData, err := usecase.GetUserById(1)
		assert.NoError(t, err)
		assert.Equal(t, resultData.ID, returnData.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Data By Id", func(t *testing.T) {
		repo.On("SelectDataById", 0).Return(user.UserCore{}, nil).Once()

		usecase := NewUserUsecase(repo)
		resultData, err := usecase.GetUserById(0)
		assert.Error(t, err)
		assert.Equal(t, resultData.ID, returnData.ID)
		repo.AssertExpectations(t)
	})
}

func TestInsertData(t *testing.T) {
	dataLayerMock := new(mocks.UserData)
	t.Run("Success Post Data", func(t *testing.T) {
		dataLayerMock.On("InsertData", mock.Anything).Return(1, nil).Once()

		dataInput := user.UserCore{Name_User: "Jono", Email: "edxample@gmail.com", Password: "qwerty", Address_user: "konohagakure"}
		usecase := NewUserUsecase(dataLayerMock)
		result, err := usecase.PostData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		dataLayerMock.AssertExpectations(t)
	})

	t.Run("failed. Name is empty", func(t *testing.T) {
		// dataLayerMock.On("InsertData", mock.Anything).Return(1, nil)
		dataInput := user.UserCore{Email: "example@gmail.com", Password: "qwerty"}
		usecase := NewUserUsecase(dataLayerMock)
		result, err := usecase.PostData(dataInput)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
		dataLayerMock.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := user.UserCore{
		ID:        1,
		Name_User: "Jono",
		Email:     "example@gmail.com",
		Password:  "qwerty",
	}
	t.Run("Success Login", func(t *testing.T) {
		repo.On("PostLogin", mock.Anything).Return(returnData, nil).Once()

		usecase := NewUserUsecase(repo)
		resultData, err := usecase.PostLogin(returnData)
		assert.NoError(t, err)
		assert.Equal(t, resultData, returnData.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Login", func(t *testing.T) {
		repo.On("PostLogin", mock.Anything).Return(user.UserCore{}, nil).Once()

		usecase := NewUserUsecase(repo)
		resultData, err := usecase.PostLogin(returnData)
		assert.Error(t, err)
		assert.Equal(t, resultData, returnData.ID)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	dataLayerMock := new(mocks.UserData)
	t.Run("Success Update Data", func(t *testing.T) {
		dataLayerMock.On("PutData", mock.Anything).Return(1, nil).Once()

		dataInput := user.UserCore{Name_User: "Jono", Email: "example@gmail.com", Password: "qwerty", Address_user: "konohagakure"}
		usecase := NewUserUsecase(dataLayerMock)
		result, err := usecase.PutData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		dataLayerMock.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Success Delete Data", func(t *testing.T) {
		dataLayerMock := new(mocks.UserData)
		dataLayerMock.On("DeleteUser", mock.Anything).Return(1, nil).Once()

		usecase := NewUserUsecase(dataLayerMock)
		result, err := usecase.DeleteUser(1)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		dataLayerMock.AssertExpectations(t)
	})

	t.Run("Failed Delete Data", func(t *testing.T) {
		dataLayerMock := new(mocks.UserData)
		dataLayerMock.On("DeleteUser", mock.Anything).Return(0, nil).Once()

		usecase := NewUserUsecase(dataLayerMock)
		result, err := usecase.DeleteUser(0)
		assert.Error(t, err)
		assert.Equal(t, 0, result)
		dataLayerMock.AssertExpectations(t)

	})
}
