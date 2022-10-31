package model

type GetCompetitionResponse struct {
	Competitions []CompetitionResponse `json:"competitions"`
}

type CompetitionResponse struct {
	ID   int `json:"id"`
	Area struct {
		ID int `json:"id"`
	} `json:"area"`
}
