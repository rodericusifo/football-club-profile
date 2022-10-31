package service_area_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/libs/util"
	dto_service_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/service/dto"
)

func (s *AreaService) GetAreas(payload *dto_service_area_v1.GetAreasDTO) (*dto_service_area_v1.AreaFromApiDTO, error) {
	modelGetAreas, err := s.AreaResource.GetAreas()
	if err != nil {
		return nil, err
	}

	areaFromApiDto, err := util.ConvertToStruct[*dto_service_area_v1.AreaFromApiDTO](modelGetAreas)
	if err != nil {
		return nil, err
	}

	areaFromApiDtoPaginated := util.PaginateSlice(areaFromApiDto.Areas, ((*payload.Query.Page)-1)*(*payload.Query.Limit), *payload.Query.Limit)

	areaFromApiDto.Meta = types.Pagination{
		Page:  *payload.Query.Page,
		Limit: *payload.Query.Limit,
	}

	areaFromApiDto.Count = len(areaFromApiDtoPaginated)
	areaFromApiDto.Areas = areaFromApiDtoPaginated

	return areaFromApiDto, nil
}
