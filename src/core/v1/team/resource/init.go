package resource_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/database"
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
	footballDataOrg_api_repository "github.com/rodericusifo/football-club-profile/src/repository/api/football-data.org"
)

type ITeamResource interface {
	GetTeamsByAreaID(query *types.Query) (*model.GetTeamsResponse, error)
	GetTeamByID(query *types.Query) (*model.GetTeamResponse, error)
}

type TeamResource struct {
	FootballDataOrgAPI *footballDataOrg_api_repository.FootballDataOrgApiRepository
}

func InitTeamResource() ITeamResource {
	var (
		footballDataOrgAPI = database.InitFootballDataOrgApi()
	)

	return &TeamResource{
		FootballDataOrgAPI: footballDataOrgAPI,
	}
}
