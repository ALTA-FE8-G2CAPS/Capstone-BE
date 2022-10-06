package usecase

import (
	"capstone-project/features/schedule"
	"capstone-project/utils/helper"
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

func (usecase *scheduleUsecase) PostData(dataSchedule schedule.ScheduleCore) (row int, err error) {
	if dataSchedule.FieldID == 0 || dataSchedule.Day == "" || dataSchedule.Start_hours == "" || dataSchedule.End_hours == "" {
		return -1, errors.New("data tidak boleh kosong")
	}

	var schedule_id int
	schedule_id, row, err = usecase.scheduleData.InsertData(dataSchedule)
	if err != nil {
		return -1, err
	}
	deteilShceduleTime := helper.GenerateSchedule(schedule_id, dataSchedule.Start_hours, dataSchedule.End_hours)
	row, err = usecase.scheduleData.InsertDetailSchedule(schedule_id, deteilShceduleTime)
	if err != nil {
		return -1, err
	}
	return row, err
}

// func (usecase *scheduleUsecase) PutData(data schedule.ScheduleCore, schedule_id int) (row int, err error){
// 	row, err = usecase.scheduleData.UpdateSchedule(data, schedule_id)
// 	if err != nil {
// 		return -1, err
// 	}
// }

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

func (usecase *scheduleUsecase) DeleteSchedule(schedule_id int) (row int, err error) {
	result, err := usecase.scheduleData.DeleteSchedule(schedule_id)
	if err != nil {
		return -1, err
	}
	return result, err
}
