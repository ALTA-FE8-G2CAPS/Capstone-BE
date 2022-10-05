package data

import (
	"capstone-project/features/schedule"

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

func (repo *scheduleData) InsertData(data schedule.ScheduleCore) (int, error) {
	newSchedule := fromCore(data)

	tx := repo.db.Create(&newSchedule)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *scheduleData) SelectAllSchedule(field_id int) ([]schedule.ScheduleCore, error) {
	var dataSchedule []Schedule

	if field_id != 0 {
		tx := repo.db.Where("field_id = ?", field_id).Preload("Field").Find(&dataSchedule)
		// fmt.Println(dataField[0].Venue.Name_venue)

		if tx.Error != nil {
			return []schedule.ScheduleCore{}, tx.Error
		}
	} else {
		tx := repo.db.Preload("Field").Find(&dataSchedule)

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
