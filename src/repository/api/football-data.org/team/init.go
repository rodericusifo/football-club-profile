package team_footballDataOrg_api_repository

import (
	"net/http"

	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
)

type ITeamRepository interface {
	GetTeamsByAreaID(query *types.Query) (*model.GetTeamsResponse, error)
	GetTeamByID(query *types.Query) (*model.GetTeamResponse, error)
}

type TeamRepository struct {
	client *http.Client
}

func InitTeamRepository(client *http.Client) ITeamRepository {
	return &TeamRepository{
		client: client,
	}
}
