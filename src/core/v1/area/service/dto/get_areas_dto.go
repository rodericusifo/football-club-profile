package dto_service_area_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/dto"
	request_controller_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/controller/request"
)

type GetAreasDTO struct {
	dto.DTO[any, *request_controller_area_v1.GetAreasRequestQuery, any]
}
