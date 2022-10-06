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
	newSchedule := fromCore(data)

	tx := repo.db.Create(&newSchedule)
	if tx.Error != nil {
		return 0, 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

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

	tx := repo.db.Where("id = ?", id).Preload("Field").First(&dataSchedule)

	if tx.Error != nil {
		return schedule.ScheduleCore{}, tx.Error
	}

	dataScheduleCore := dataSchedule.toCore()
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

// func (repo *scheduleData) UpdateSchedule(data schedule.ScheduleCore, schedule_id int) (int, int, error) {
// 	rowDel, errDel := repo.DeleteSchedule(schedule_id)
// 	if errDel != nil {
// 		return 0, 0, errDel
// 	}
// 	return rowDel, 0, nil

// 	// schedule_id, row, errPost := repo.InsertData(data)
// 	// if errPost != nil {
// 	// 	return 0, 0, errPost
// 	// }
// 	// return schedule_id, row, nil
// }
