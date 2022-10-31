package resource_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *TeamResource) GetTeamByID(query *types.Query) (*model.GetTeamResponse, error) {
	return r.FootballDataOrgAPI.TeamRepository.GetTeamByID(query)
}
