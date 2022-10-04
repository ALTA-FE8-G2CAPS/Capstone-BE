package usecase

import (
	"capstone-project/features/field"
	"errors"
)

type FieldUsecase struct {
	fieldData field.DataInterface
}

func New(data field.DataInterface) field.UsecaseInterface {
	return &FieldUsecase{
		data, // fieldData
	}
}

func (usecase *FieldUsecase) GetAllField(venue_id int) ([]field.FieldCore, error) {
	dataField, err := usecase.fieldData.SelectAllField(venue_id)
	if err != nil {
		return nil, err
	}
	return dataField, nil
}

func (usecase *FieldUsecase) PostData(data field.FieldCore) (row int, err error) {
	if data.VenueID == 0 || data.Category == "" || data.Price == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err = usecase.fieldData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *FieldUsecase) GetFieldById(id int) (field.FieldCore, error) {
	result, err := usecase.fieldData.SelectFieldById(id)
	if err != nil {
		return field.FieldCore{}, errors.New("data tidak ditemukan")
	}
	return result, nil
}

func (usecase *FieldUsecase) DeleteField(user_id, venue_id, field_id int) (row int, err error) {
	result, err := usecase.fieldData.DeleteField(user_id, venue_id, field_id)
	if err != nil {
		return -1, err
	}
	return result, err
}

func (usecase *FieldUsecase) PutData(data field.FieldCore, user_id, venue_id, field_id int) (row int, err error) {
	row, err = usecase.fieldData.UpdateField(data, user_id, venue_id, field_id)
	if err != nil {
		return -1, err
	}
	return row, err
}
