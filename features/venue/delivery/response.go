package delivery

import "capstone-project/features/venue"

type VenueResponse struct {
	ID                uint    `json:"id" form:"id"`
	Name_venue        string  `json:"name_venue" form:"name_venue"`
	Address_venue     string  `json:"address_venue" form:"address_venue"`
	Description_venue string  `json:"description_venue" form:"description_venue"`
<<<<<<< HEAD
	Nama_user         string  `json:"nama_user" form:"nama_user"`
=======
	Name_user         string  `json:"name_user" form:"name_user"`
>>>>>>> 31e7c5130e5b0e287b19439499d3cbc57a4d8dff
	Latitude          float64 `json:"latitude" form:"latitude"`
	Longitude         float64 `json:"longitude" form:"longitude"`
}

func FromCore(data venue.VenueCore) VenueResponse {
	return VenueResponse{
		ID:                data.ID,
		Name_venue:        data.Name_venue,
		Address_venue:     data.Address_venue,
		Description_venue: data.Description_venue,
		Name_user:         data.Name_user,
		Latitude:          data.Latitude,
		Longitude:         data.Longitude,
	}

}

func FromCoreList(data []venue.VenueCore) []VenueResponse {
	var list []VenueResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
