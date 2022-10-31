package service_team_v1

import (
	resource_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/resource"
	dto_service_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/service/dto"
)

type ITeamService interface {
	GetTeamsByAreaID(payload *dto_service_team_v1.GetTeamsByAreaIDDTO) (*dto_service_team_v1.TeamFromApiDTO, error)
	GetTeamByID(payload *dto_service_team_v1.GetTeamByIDDTO) (*dto_service_team_v1.TeamDetailFromApiDTO, error)
}

type TeamService struct {
	TeamResource resource_team_v1.ITeamResource
}

func InitTeamService(teamResource resource_team_v1.ITeamResource) ITeamService {
	return &TeamService{
		TeamResource: teamResource,
	}
}
