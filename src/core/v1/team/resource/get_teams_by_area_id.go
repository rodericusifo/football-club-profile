package resource_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *TeamResource) GetTeamsByAreaID(query *types.Query) (*model.GetTeamsResponse, error) {
	return r.FootballDataOrgAPI.TeamRepository.GetTeamsByAreaID(query)
}
