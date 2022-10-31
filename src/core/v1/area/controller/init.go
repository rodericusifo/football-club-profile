package controller_area_v1

import (
	"github.com/labstack/echo/v4"
	service_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/service"
)

type AreaController struct {
	AreaService service_area_v1.IAreaService
}

func InitAreaController(areaService service_area_v1.IAreaService) *AreaController {
	return &AreaController{AreaService: areaService}
}

func (areaController *AreaController) Mount(group *echo.Group) {
	group.GET("", areaController.GetAreas)
}
