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

func (c *TeamController) GetTeamByID(ctx echo.Context) error {
	reqParam := new(request_controller_team_v1.GetTeamByIDRequestParam)
	if err := validator.ValidateRequest(ctx, reqParam); err != nil {
		return err
	}

	teamFromApiDto, err := c.TeamService.GetTeamByID(&dto_service_team_v1.GetTeamByIDDTO{
		DTO: dto.DTO[*request_controller_team_v1.GetTeamByIDRequestParam, any, any]{
			Param: reqParam,
			Query: nil,
			Body:  nil,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("get detail team by id success", teamFromApiDto))
}
