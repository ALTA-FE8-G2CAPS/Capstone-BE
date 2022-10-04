package data

import (
	"capstone-project/features/field"
	"fmt"

	"gorm.io/gorm"
)

type fieldData struct {
	db *gorm.DB
}

func New(db *gorm.DB) field.DataInterface {
	return &fieldData{
		db: db,
	}

}

func (repo *fieldData) InsertData(data field.FieldCore) (int, error) {
	newField := fromCore(data)

	tx := repo.db.Create(&newField)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// token, errToken := middlewares.CreateToken(int(newUser.ID))
	// if errToken != nil {
	// 	return "", -1, errToken
	// }

	return int(tx.RowsAffected), nil
}

func (repo *fieldData) SelectAllField(user_id int) ([]field.FieldCore, error) {
	var dataField []Field

	if user_id != 0 {
		tx := repo.db.Where("user_id = ?", user_id).Preload("Venue").Find(&dataField)
		fmt.Println(dataField[0].Venue.Name_venue)

		if tx.Error != nil {
			return []field.FieldCore{}, tx.Error
		}
	} else {
		tx := repo.db.Preload("Venue").Find(&dataField)

		if tx.Error != nil {
			return []field.FieldCore{}, tx.Error
		}
	}
	return toCoreList(dataField), nil

}

func (repo *fieldData) SelectFieldById(id int) (field.FieldCore, error) {
	var dataField Field
	dataField.ID = uint(id)

	tx := repo.db.Where("id = ?", id).Preload("Venue").First(&dataField)

	if tx.Error != nil {
		return field.FieldCore{}, tx.Error
	}

	dataFieldCore := dataField.toCore()
	return dataFieldCore, nil

}
