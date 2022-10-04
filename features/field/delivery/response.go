package delivery

import (
	"capstone-project/features/field"
)

type FieldResponse struct {
	ID         uint   `json:"id"`
	VenueID    uint   `json:"venue_id"`
	Name_venue string `json:"name_venue"`
	Category   string `json:"category"`
	Price      uint   `json:"price"`
}

func FromCore(data field.FieldCore) FieldResponse {
	return FieldResponse{
		ID:         data.ID,
		VenueID:    data.VenueID,
		Name_venue: data.Name_venue,
		Category:   data.Category,
		Price:      data.Price,
	}

}

func FromCoreList(data []field.FieldCore) []FieldResponse {
	var list []FieldResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
