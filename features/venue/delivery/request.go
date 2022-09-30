package delivery

import "capstone-project/features/venue"

type VenueRequest struct {
	Name_venue        string  `json:"name_venue" form:"name_venue"`
	Address_venue     string  `json:"address_venue" form:"address_venue"`
	UserID            uint    `json:"user_id" form:"user_id"`
	Latitude          float64 `json:"latitude" form:"latitude"`
	Longitude         float64 `json:"longitude" form:"longitude"`
	Description_venue string  `json:"description_venue" form:"description_venue"`
}

type Foto_venueRequest struct {
	VenueID    uint   `json:"venue_id" form:"venue_id"`
	Foto_venue string `json:"foto_venue" form:"foto_venue"`
}

func ToCoreFoto_venue(data Foto_venueRequest) venue.FotoVenue {
	return venue.FotoVenue{
		VenueID:    data.VenueID,
		Foto_Venue: data.Foto_venue,
	}
}

func ToCore(data VenueRequest) venue.VenueCore {
	return venue.VenueCore{
		Name_Venue:        data.Name_venue,
		Address_Venue:     data.Address_venue,
		UserID:            data.UserID,
		Latitude:          data.Latitude,
		Longitude:         data.Longitude,
		Description_venue: data.Description_venue,
	}
}
