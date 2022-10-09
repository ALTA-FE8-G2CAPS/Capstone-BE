package delivery

import "capstone-project/features/venue"

type VenueResponse struct {
	ID                uint        `json:"id" form:"id"`
	Name_venue        string      `json:"name_venue" form:"name_venue"`
	Address_venue     string      `json:"address_venue" form:"address_venue"`
	Description_venue string      `json:"description_venue" form:"description_venue"`
	UserID            uint        `json:"user_id" form:"user_id"`
	Name_user         string      `json:"name_user" form:"name_user"`
	Latitude          float64     `json:"latitude" form:"latitude"`
	Longitude         float64     `json:"longitude" form:"longitude"`
	Max_price         uint        `json:"max_price"`
	Min_price         uint        `json:"min_price"`
	Foto_venue        []FotoVenue `json:"foto_venue"`
}

type FotoVenue struct {
	ID         uint   `json:"foto_venue_id"`
	VenueID    uint   `json:"venue_id"`
	Foto_Venue string `json:"foto_venue"`
}

type Field struct {
	Min uint `json:"min_price"`
	Max uint `json:"max_price"`
}

func FromCore(data venue.VenueCore) VenueResponse {
	return VenueResponse{
		ID:                data.ID,
		Name_venue:        data.Name_venue,
		Address_venue:     data.Address_venue,
		Description_venue: data.Description_venue,
		UserID:            data.UserID,
		Name_user:         data.Name_user,
		Latitude:          data.Latitude,
		Longitude:         data.Longitude,
		Max_price:         data.Max_price,
		Min_price:         data.Min_price,
		Foto_venue:        FromCoreFotoList(data.Foto_venue),
	}

}

func FromCoreList(data []venue.VenueCore) []VenueResponse {
	var list []VenueResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}

func FromCoreFotoList(data []venue.FotoVenue) []FotoVenue {
	var list []FotoVenue
	for _, v := range data {
		var foto_venue FotoVenue
		foto_venue.ID = v.ID
		foto_venue.VenueID = v.VenueID
		foto_venue.Foto_Venue = v.Foto_Venue
		list = append(list, foto_venue)
	}
	return list
}
