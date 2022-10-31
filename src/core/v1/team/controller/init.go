package controller_team_v1

import (
	"github.com/labstack/echo/v4"
	service_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/service"
)

type TeamController struct {
	TeamService service_team_v1.ITeamService
}

func InitTeamController(teamService service_team_v1.ITeamService) *TeamController {
	return &TeamController{TeamService: teamService}
}

func (teamController *TeamController) Mount(group *echo.Group) {
	group.GET("", teamController.GetTeamsByAreaID)
	group.GET("/:id", teamController.GetTeamByID)
}
