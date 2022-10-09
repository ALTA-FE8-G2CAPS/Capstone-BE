package usecase

import (
	"capstone-project/features/field"
	mocks "capstone-project/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllField(t *testing.T) {
	repo := new(mocks.FieldData)
	returnData := []field.FieldCore{{ID: 1, VenueID: 1, Name_venue: "gor bagus coba", Category: "fursal", Price: 10000, Schedule: field.Schedule{ID: 1, FieldID: 1, Day: "senin", Start_hours: 1, End_hours: 2, SchedulesDetail: []field.ScheduleDetail{{ID: 1, ScheduleID: 1, Start_hours: 1, End_hours: 2, Status_schedule: "available"}}}}}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("SelectAllField", 1).Return(returnData, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetAllField(1)
		assert.NoError(t, err)
		assert.Equal(t, resultData[0].ID, returnData[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All Data", func(t *testing.T) {
		repo.On("SelectAllField", 0).Return([]field.FieldCore{}, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetAllField(0)
		assert.Error(t, err)
		assert.Equal(t, resultData[0].ID, 0)
		repo.AssertExpectations(t)
	})
}

func TestGetFieldById(t *testing.T) {
	repo := new(mocks.FieldData)
	returnData := field.FieldCore{
		ID:         1,
		VenueID:    1,
		Name_venue: "gor bagus coba",
		Category:   "fursal",
		Price:      10000,
		Schedule: field.Schedule{
			ID:          1,
			FieldID:     1,
			Day:         "senin",
			Start_hours: 1,
			End_hours:   2,
			SchedulesDetail: []field.ScheduleDetail{
				{ID: 1, ScheduleID: 1, Start_hours: 1, End_hours: 2, Status_schedule: "available"},
			},
		},
	}

	t.Run("Success Get Data By Id", func(t *testing.T) {
		repo.On("SelectFieldById", 1).Return(returnData, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetFieldById(1)
		assert.NoError(t, err)
		assert.Equal(t, resultData.ID, returnData.ID)
		repo.AssertExpectations(t)
	})
}

func TestPostData(t *testing.T) {
	repo := new(mocks.FieldData)
	returnData := field.FieldCore{
		ID:         1,
		VenueID:    1,
		Name_venue: "gor bagus coba",
		Category:   "fursal",
		Price:      10000,
		Schedule: field.Schedule{
			ID:          1,
			FieldID:     1,
			Day:         "senin",
			Start_hours: 1,
			End_hours:   2,
			SchedulesDetail: []field.ScheduleDetail{
				{ID: 1, ScheduleID: 1, Start_hours: 1, End_hours: 2, Status_schedule: "available"},
			},
		},
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
	repo := new(mocks.FieldData)
	returnData := field.FieldCore{
		ID:         1,
		VenueID:    1,
		Name_venue: "gor bagus coba",
		Category:   "fursal",
		Price:      10000,
		Schedule: field.Schedule{
			ID:          1,
			FieldID:     1,
			Day:         "senin",
			Start_hours: 1,
			End_hours:   2,
			SchedulesDetail: []field.ScheduleDetail{
				{ID: 1, ScheduleID: 1, Start_hours: 1, End_hours: 2, Status_schedule: "available"},
			},
		},
	}

	t.Run("Success Put Data", func(t *testing.T) {
		repo.On("UpdateData", returnData).Return(1, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.PutData(returnData, int(returnData.ID))
		assert.NoError(t, err)
		assert.Equal(t, resultData, 1)
		repo.AssertExpectations(t)
	})
}

func TestDeleteField(t *testing.T) {
	repo := new(mocks.FieldData)
	returnData := field.FieldCore{
		ID:         1,
		VenueID:    1,
		Name_venue: "gor bagus coba",
		Category:   "fursal",
		Price:      10000,
		Schedule: field.Schedule{
			ID:          1,
			FieldID:     1,
			Day:         "senin",
			Start_hours: 1,
			End_hours:   2,
			SchedulesDetail: []field.ScheduleDetail{
				{ID: 1, ScheduleID: 1, Start_hours: 1, End_hours: 2, Status_schedule: "available"},
			},
		},
	}

	t.Run("Success Delete Data", func(t *testing.T) {
		repo.On("DeleteData", returnData).Return(1, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.DeleteField(int(returnData.ID))
		assert.NoError(t, err)
		assert.Equal(t, resultData, 1)
		repo.AssertExpectations(t)
	})
}
