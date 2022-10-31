package footballDataOrg_api_repository

import (
	"net/http"

	area_footballDataOrg_api_repository "github.com/rodericusifo/football-club-profile/src/repository/api/football-data.org/area"
	team_footballDataOrg_api_repository "github.com/rodericusifo/football-club-profile/src/repository/api/football-data.org/team"
)

type FootballDataOrgApiRepository struct {
	AreaRepository area_footballDataOrg_api_repository.IAreaRepository
	TeamRepository team_footballDataOrg_api_repository.ITeamRepository
}

func InitFootballDataOrgApiRepository(client *http.Client) *FootballDataOrgApiRepository {
	var (
		areaRepository = area_footballDataOrg_api_repository.InitAreaRepository(client)
		teamRepository = team_footballDataOrg_api_repository.InitTeamRepository(client)
	)

	return &FootballDataOrgApiRepository{
		AreaRepository: areaRepository,
		TeamRepository: teamRepository,
	}
}
