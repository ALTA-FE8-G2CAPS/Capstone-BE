package usecase

import (
	"capstone-project/features/field"
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
	dataField, err := usecase.fieldData.SelectAllField()
	if err != nil {
		return nil, err
	}
	return dataField, nil
}
