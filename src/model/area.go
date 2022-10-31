package model

type GetAreasResponse struct {
	Count   int                `json:"count"`
	Filters FilterAreaResponse `json:"filters"`
	Areas   []AreaResponse     `json:"areas"`
}

type FilterAreaResponse struct{}

type AreaResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"countryCode"`
	Flag         *string `json:"flag"`
	ParentAreaID int     `json:"parentAreaId"`
	ParentArea   string  `json:"parentArea"`
}
