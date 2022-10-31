package dto_service_team_v1

import "github.com/rodericusifo/football-club-profile/libs/types"

type TeamFromApiDTO struct {
	Meta    types.Pagination `json:"meta"`
	Count   int              `json:"count"`
	Filters FilterFromApi    `json:"filters"`
	Teams   []TeamFromAPI    `json:"teams"`
}

type FilterFromApi struct{}

type TeamFromAPI struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	ShortName  string          `json:"shortName"`
	Tla        string          `json:"tla"`
	Crest      string          `json:"crest"`
	Website    string          `json:"website"`
	Founded    int             `json:"founded"`
	ClubColors string          `json:"clubColors"`
	Venue      string          `json:"venue"`
	Area       TeamAreaFromAPI `json:"area"`
}

type TeamAreaFromAPI struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type TeamDetailFromApiDTO struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	ShortName  string          `json:"shortName"`
	Tla        string          `json:"tla"`
	Crest      string          `json:"crest"`
	Website    string          `json:"website"`
	Founded    int             `json:"founded"`
	ClubColors string          `json:"clubColors"`
	Venue      string          `json:"venue"`
	Area       TeamAreaFromAPI `json:"area"`
}