package service_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/libs/util"
	dto_service_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/service/dto"
)

func (s *TeamService) GetTeamByID(payload *dto_service_team_v1.GetTeamByIDDTO) (*dto_service_team_v1.TeamDetailFromApiDTO, error) {
	modelGetTeam, err := s.TeamResource.GetTeamByID(&types.Query{
		QueryTeamByID: types.QueryTeamByID{
			TeamID: payload.Param.TeamID,
		},
	})
	if err != nil {
		return nil, err
	}

	teamDetailFromApiDto, err := util.ConvertToStruct[*dto_service_team_v1.TeamDetailFromApiDTO](modelGetTeam)
	if err != nil {
		return nil, err
	}

	return teamDetailFromApiDto, nil
}
