package controller_area_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/football-club-profile/libs/dto"
	"github.com/rodericusifo/football-club-profile/libs/response"
	"github.com/rodericusifo/football-club-profile/libs/validator"
	request_controller_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/controller/request"
	dto_service_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/service/dto"
)

func (c *AreaController) GetAreas(ctx echo.Context) error {
	reqQuery := new(request_controller_area_v1.GetAreasRequestQuery)
	if err := validator.ValidateRequest(ctx, reqQuery); err != nil {
		return err
	}

	areaFromApiDto, err := c.AreaService.GetAreas(&dto_service_area_v1.GetAreasDTO{
		DTO: dto.DTO[any, *request_controller_area_v1.GetAreasRequestQuery, any]{
			Param: nil,
			Query: reqQuery,
			Body:  nil,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("get list area success", areaFromApiDto))
}
