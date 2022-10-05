package usecase

import (
	"capstone-project/features/schedule"
	"errors"
)

type scheduleUsecase struct {
	scheduleData schedule.DataInterface
}

func New(data schedule.DataInterface) schedule.UsecaseInterface {
	return &scheduleUsecase{
		data,
	}
}

func (usecase *scheduleUsecase) PostData(data schedule.ScheduleCore) (row int, err error) {
	if data.FieldID == 0 || data.Category == "" || data.Day == "" || data.Start_hours == "" || data.End_hours == "" {
		return -1, errors.New("data tidak boleh kosong")
	}
	row, err = usecase.scheduleData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *scheduleUsecase) GetAllSchedule(field_id int) ([]schedule.ScheduleCore, error) {
	dataMentee, err := usecase.scheduleData.SelectAllSchedule(field_id)
	return dataMentee, err

}

func (usecase *scheduleUsecase) GetScheduleById(id int) (schedule.ScheduleCore, error) {
	result, err := usecase.scheduleData.SelectScheduleById(id)
	if err != nil {
		return schedule.ScheduleCore{}, err
	}
	return result, nil
}
