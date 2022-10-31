package resource_area_v1

import "github.com/rodericusifo/football-club-profile/src/model"

func (r *AreaResource) GetAreas() (*model.GetAreasResponse, error) {
	return r.FootballDataOrgAPI.AreaRepository.GetAreas()
}
