package service_area_v1

import (
	resource_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/resource"
	dto_service_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/service/dto"
)

type IAreaService interface {
	GetAreas(payload *dto_service_area_v1.GetAreasDTO) (*dto_service_area_v1.AreaFromApiDTO, error)
}

type AreaService struct {
	AreaResource resource_area_v1.IAreaResource
}

func InitAreaService(areaResource resource_area_v1.IAreaResource) IAreaService {
	return &AreaService{
		AreaResource: areaResource,
	}
}
