package types

type Query struct {
	SelectColumns []string
	QueryTeamsByAreaID
	QueryTeamByID
}

type QueryTeamsByAreaID struct {
	AreaIDs []int
}

type QueryTeamByID struct {
	TeamID int
}
