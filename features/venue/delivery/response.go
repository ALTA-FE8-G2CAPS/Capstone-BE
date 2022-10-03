package delivery

import "capstone-project/features/venue"

type VenueResponse struct {
	ID                uint    `json:"id" form:"id"`
	Name_venue        string  `json:"name_venue" form:"name_venue"`
	Address_venue     string  `json:"address_venue" form:"address_venue"`
	Description_venue string  `json:"description_venue" form:"description_venue"`
	UserID            uint    `json:"user_id" form:"user_id"`
	Name_user         string  `json:"name_user" form:"name_user"`
	Latitude          float64 `json:"latitude" form:"latitude"`
	Longitude         float64 `json:"longitude" form:"longitude"`
	Foto              []Foto_venueResponse
}

type Foto_venueResponse struct {
	Foto_Venue string `json:"foto_venue" form:"foto_venue"`
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
	}

}

func FromCoreFoto(data venue.FotoVenue) Foto_venueResponse {
	return Foto_venueResponse{
		Foto_Venue: data.Foto_Venue,
	}
}

func FromCoreFotoList(data []venue.FotoVenue) []Foto_venueResponse {
	var list []Foto_venueResponse
	for _, v := range data {
		list = append(list, FromCoreFoto(v))
	}
	return list
}
func FromCoreList(data []venue.VenueCore) []VenueResponse {
	var list []VenueResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
