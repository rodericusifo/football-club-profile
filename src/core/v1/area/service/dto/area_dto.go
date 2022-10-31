package dto_service_area_v1

import "github.com/rodericusifo/football-club-profile/libs/types"

type AreaFromApiDTO struct {
	Meta    types.Pagination `json:"meta"`
	Count   int              `json:"count"`
	Filters FilterFromApi    `json:"filters"`
	Areas   []AreaFromAPI    `json:"areas"`
}

type FilterFromApi struct{}

type AreaFromAPI struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"countryCode"`
	Flag         *string `json:"flag"`
	ParentAreaID int     `json:"parentAreaId"`
	ParentArea   string  `json:"parentArea"`
}
