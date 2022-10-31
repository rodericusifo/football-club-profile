package team_footballDataOrg_api_repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodericusifo/football-club-profile/libs/config"
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/libs/util"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *TeamRepository) GetTeamsByAreaID(query *types.Query) (*model.GetTeamsResponse, error) {
	var (
		err                 error
		dataCompetetion     model.GetCompetitionResponse
		dataTeam, dataTeams model.GetTeamsResponse
	)

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/competitions", config.AppConfig.API.FootballDataOrg.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Auth-Token", config.AppConfig.API.FootballDataOrg.AuthToken)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&dataCompetetion)
	if err != nil {
		return nil, err
	}

	competitionIDs := []int{}

	if len(query.AreaIDs) > 0 {
		for _, competition := range dataCompetetion.Competitions {
			if util.SliceContains(query.AreaIDs, competition.Area.ID) {
				competitionIDs = append(competitionIDs, competition.ID)
			}
		}
	} else {
		for _, competition := range dataCompetetion.Competitions {
			competitionIDs = append(competitionIDs, competition.ID)
		}
	}

	for _, competitionID := range competitionIDs {
		request, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/competitions/%d/teams", config.AppConfig.API.FootballDataOrg.BaseUrl, competitionID), nil)
		if err != nil {
			return nil, err
		}
		request.Header.Set("X-Auth-Token", config.AppConfig.API.FootballDataOrg.AuthToken)

		response, err := r.client.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(&dataTeam)
		if err != nil {
			return nil, err
		}

		dataTeams.Count += len(dataTeam.Teams)
		dataTeams.Teams = append(dataTeams.Teams, dataTeam.Teams...)
	}

	return &dataTeams, nil
}
