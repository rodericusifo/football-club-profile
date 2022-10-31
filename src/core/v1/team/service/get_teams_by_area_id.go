package service_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/libs/util"
	dto_service_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/service/dto"
)

func (s *TeamService) GetTeamsByAreaID(payload *dto_service_team_v1.GetTeamsByAreaIDDTO) (*dto_service_team_v1.TeamFromApiDTO, error) {
	modelGetTeams, err := s.TeamResource.GetTeamsByAreaID(&types.Query{
		QueryTeamsByAreaID: types.QueryTeamsByAreaID{
			AreaIDs: payload.Query.AreaIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	teamFromApiDto, err := util.ConvertToStruct[*dto_service_team_v1.TeamFromApiDTO](modelGetTeams)
	if err != nil {
		return nil, err
	}

	teamFromApiDto.Teams = util.UniqueSlice(teamFromApiDto.Teams)

	teamFromApiDtoPaginated := util.PaginateSlice(teamFromApiDto.Teams, ((*payload.Query.Page)-1)*(*payload.Query.Limit), *payload.Query.Limit)
	
	teamFromApiDto.Meta = types.Pagination{
		Page:  *payload.Query.Page,
		Limit: *payload.Query.Limit,
	}

	teamFromApiDto.Count = len(teamFromApiDtoPaginated)
	teamFromApiDto.Teams = teamFromApiDtoPaginated

	return teamFromApiDto, nil
}
