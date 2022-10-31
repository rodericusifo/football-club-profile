package controller_team_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/football-club-profile/libs/dto"
	"github.com/rodericusifo/football-club-profile/libs/response"
	"github.com/rodericusifo/football-club-profile/libs/validator"
	request_controller_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/controller/request"
	dto_service_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/service/dto"
)

func (c *TeamController) GetTeamsByAreaID(ctx echo.Context) error {
	reqQuery := new(request_controller_team_v1.GetTeamsByAreaIDRequestQuery)
	if err := validator.ValidateRequest(ctx, reqQuery); err != nil {
		return err
	}

	teamFromApiDto, err := c.TeamService.GetTeamsByAreaID(&dto_service_team_v1.GetTeamsByAreaIDDTO{
		DTO: dto.DTO[any, *request_controller_team_v1.GetTeamsByAreaIDRequestQuery, any]{
			Param: nil,
			Query: reqQuery,
			Body:  nil,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("get list team by area ids success", teamFromApiDto))
}
