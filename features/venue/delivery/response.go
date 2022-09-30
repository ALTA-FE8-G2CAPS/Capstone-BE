package delivery

import "capstone-project/features/venue"

type VenueResponse struct {
	ID                uint
	Name_venue        string
	Address_venue     string
	Description_venue string
	Nama_user         string
	Latitude          float64
	Longitude         float64
}

func FromCore(data venue.VenueCore) VenueResponse {
	return VenueResponse{
		ID:                data.ID,
		Name_venue:        data.Name_venue,
		Address_venue:     data.Address_venue,
		Description_venue: data.Description_venue,
		Nama_user:         data.Nama_user,
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
