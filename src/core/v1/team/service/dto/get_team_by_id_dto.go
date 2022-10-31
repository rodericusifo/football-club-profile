package dto_service_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/dto"
	request_controller_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/controller/request"
)

type GetTeamByIDDTO struct {
	dto.DTO[*request_controller_team_v1.GetTeamByIDRequestParam, any, any]
}
