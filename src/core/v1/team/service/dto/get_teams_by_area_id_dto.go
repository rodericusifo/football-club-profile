package dto_service_team_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/dto"
	request_controller_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/controller/request"
)

type GetTeamsByAreaIDDTO struct {
	dto.DTO[any, *request_controller_team_v1.GetTeamsByAreaIDRequestQuery, any]
}
