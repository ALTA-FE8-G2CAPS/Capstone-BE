package delivery

import "capstone-project/features/field"

type FieldRequest struct {
	VenueID  uint   `json:"venue_id" form:"venue_id"`
	Category string `json:"category" form:"category"`
	Price    uint   `json:"price" form:"price"`
}

func ToCore(data FieldRequest) field.FieldCore {
	return field.FieldCore{
		VenueID:  data.VenueID,
		Category: data.Category,
		Price:    data.Price,
	}

}
