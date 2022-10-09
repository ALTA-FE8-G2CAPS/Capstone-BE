package data

import (
	"capstone-project/features/schedule"
	"errors"

	"gorm.io/gorm"
)

type scheduleData struct {
	db *gorm.DB
}

func New(db *gorm.DB) schedule.DataInterface {
	return &scheduleData{
		db: db,
	}

}

func (repo *scheduleData) InsertData(data schedule.ScheduleCore) (int, int, error) {
	var Exist int64
	txUnique := repo.db.Model(&Schedule{}).Where("field_id = ?", data.FieldID).Where("day = ?", data.Day).Count(&Exist)
	if txUnique.Error != nil {
		if !errors.Is(txUnique.Error, gorm.ErrRecordNotFound) {
			return 0, 0, txUnique.Error
		}
	}
	if Exist > 0 {
		return 0, 0, errors.New("hari sudah ditambahkan")
	}
	newSchedule := fromCore(data)

	tx := repo.db.Create(&newSchedule)
	if tx.Error != nil {
		return 0, 0, tx.Error
	}
	return int(newSchedule.ID), int(tx.RowsAffected), nil
}

func (repo *scheduleData) InsertDetailSchedule(schedule_id int, dataGenerete []map[string]interface{}) (int, error) {
	tx := repo.db.Model(&ScheduleDetail{}).Create(dataGenerete)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *scheduleData) SelectAllSchedule(field_id int) ([]schedule.ScheduleCore, error) {
	var dataSchedule []Schedule

	if field_id != 0 {
		tx := repo.db.Where("field_id = ?", field_id).Preload("Field").Preload("ScheduleDetails").Find(&dataSchedule)
		// fmt.Println(dataSchedule[0].ScheduleDetails[0].Start_hours)

		if tx.Error != nil {
			return []schedule.ScheduleCore{}, tx.Error
		}
	} else {
		tx := repo.db.Preload("Field").Preload("ScheduleDetails").Find(&dataSchedule)

		if tx.Error != nil {
			return []schedule.ScheduleCore{}, tx.Error
		}
	}
	return toCoreList(dataSchedule), nil

}

func (repo *scheduleData) SelectScheduleById(id int) (schedule.ScheduleCore, error) {
	var dataSchedule Schedule
	dataSchedule.ID = uint(id)

	tx := repo.db.Where("id = ?", id).Preload("Field").Preload("ScheduleDetails").First(&dataSchedule)

	if tx.Error != nil {
		return schedule.ScheduleCore{}, tx.Error
	}

	dataScheduleCore := dataSchedule.toCore()
	return dataScheduleCore, nil

}
func (repo *scheduleData) SelectScheduleDetailById(id int) (schedule.ScheduleDetailCore, error) {
	var dataSchedule ScheduleDetail
	dataSchedule.ID = uint(id)

	tx := repo.db.Where("id = ?", id).Preload("Schedule.Field").First(&dataSchedule)

	if tx.Error != nil {
		return schedule.ScheduleDetailCore{}, tx.Error
	}

	dataScheduleCore := dataSchedule.toCoreScheduleDetail()

	return dataScheduleCore, nil

}

func (repo *scheduleData) DeleteSchedule(schedule_id int) (int, error) {
	var dataSchedule Schedule

	tx := repo.db.Where("id = ?", schedule_id).Unscoped().Delete(&dataSchedule)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed, rows affected 0")
	}
	return int(tx.RowsAffected), nil
}

func (repo *scheduleData) UpdateScheduleDetail(data schedule.ScheduleDetailCore, id int) (int, error) {
	var scheduleDetailUpdate ScheduleDetail
	txDataOld := repo.db.First(&scheduleDetailUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.ScheduleID != 0 {
		scheduleDetailUpdate.ScheduleID = data.ScheduleID
	}

	if data.Start_hours != "" {
		scheduleDetailUpdate.Start_hours = data.Start_hours
	}

	if data.End_hours != "" {
		scheduleDetailUpdate.End_hours = data.End_hours
	}

	if data.Status_schedule != "" {
		scheduleDetailUpdate.Status_schedule = data.Status_schedule
	}

	tx := repo.db.Save(&scheduleDetailUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *scheduleData) UpdateSchedule(data schedule.ScheduleCore, id int) (int, error) {
	tx := repo.db.Model(&ScheduleDetail{}).Where("schedule_id = ?", id).Select("schedule_id").Update("schedule_id", -1)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
