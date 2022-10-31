package model

type GetTeamsResponse struct {
	Count   int                `json:"count"`
	Filters FilterTeamResponse `json:"filters"`
	Teams   []TeamResponse     `json:"teams"`
}

type FilterTeamResponse struct {
	Season *string `json:"season,omitempty"`
}

type TeamResponse struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	ShortName  string           `json:"shortName"`
	Tla        string           `json:"tla"`
	Crest      string           `json:"crest"`
	Website    string           `json:"website"`
	Founded    int              `json:"founded"`
	ClubColors string           `json:"clubColors"`
	Venue      string           `json:"venue"`
	Area       TeamAreaResponse `json:"area"`
}

type TeamAreaResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type GetTeamResponse struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	ShortName  string           `json:"shortName"`
	Tla        string           `json:"tla"`
	Crest      string           `json:"crest"`
	Website    string           `json:"website"`
	Founded    int              `json:"founded"`
	ClubColors string           `json:"clubColors"`
	Venue      string           `json:"venue"`
	Area       TeamAreaResponse `json:"area"`
}